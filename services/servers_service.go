package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/Encedeus/panel/config"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/ent/node"
	"github.com/Encedeus/panel/ent/server"
	"github.com/Encedeus/panel/ent/user"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
	"github.com/Encedeus/panel/proto/go"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/google/uuid"
	"io"
	"log"
	"math"
)

func CreateServer(ctx context.Context, db *ent.Client, store *module.Store, req *protoapi.ServersCreateRequest) (*protoapi.ServersCreateResponse, error) {
	owner, _ := FindOneUser(ctx, db, &protoapi.UserFindOneRequest{UserId: req.Owner})
	if owner == nil {
		return nil, ErrUserNotFound
	}

	_, err := FindServerByNameAndOwner(ctx, db, req.Name, proto.ProtoUUIDToUUID(req.Owner))
	if err == nil {
		return nil, ErrServerAlreadyExists
	}

	/*	_, err = db.Node.Query().Where(node.IDEQ(proto.ProtoUUIDToUUID(req.Node))).First(ctx)
		if err != nil {
			return nil, ErrNodeNotFound
		}*/

	n, err := FindOptimalNode(ctx, db, req)
	if err != nil {
		return nil, err
	}
	req.Node = proto.UUIDToProtoUUID(n.ID)

	var client struct {
		CreateServer func(req *protoapi.ServersCreateRequest, id string) (*protoapi.ServersCreateResponse, error)
	}

	variant := FindStoreCraterVariant(store, req.Crater, req.CraterVariant)
	if variant == nil {
		return nil, ErrUnsupportedVariant
	}

	closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("http://localhost:%v", variant.Crater.Provider.Backend.RPCPort), "CratersHandler", &client, nil)
	if err != nil {
		return nil, err
	}
	defer closer()

	serverId := uuid.New()

	resp, err := client.CreateServer(req, serverId.String())
	if err != nil {
		return nil, err
	}

	_, err = db.Server.Create().
		SetID(serverId).
		SetName(req.Name).
		SetCraterProvider(variant.Crater.Provider.Manifest.Name).
		SetCrater(req.Crater).
		SetCraterVariant(req.CraterVariant).
		SetOwnerID(proto.ProtoUUIDToUUID(req.Owner)).
		SetRAM(req.Ram).
		SetStorage(req.Storage).
		SetLogicalCores(uint(req.LogicalCores)).
		SetNode(n).
		SetContainerId(resp.Servers[0].ContainerId).
		SetPort(uint(uint16(resp.Servers[0].Port.Value))).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FindStoreCraterVariant(store *module.Store, crater string, variant string) *module.Variant {
	for _, m := range store.Modules {
		for _, c := range m.Manifest.Backend.Craters {
			if c.Id != crater {
				continue
			}
			if c.Provider == nil {
				c.Provider = m
			}
			for _, v := range c.Variants {
				if v.Id != variant {
					continue
				}
				if v.Crater == nil {
					v.Crater = c
				}
				return v
			}
		}
	}

	return nil
}

func FindServerByNameAndOwner(ctx context.Context, db *ent.Client, name string, owner uuid.UUID) (*ent.Server, error) {
	srv, err := db.Server.Query().Where(server.HasOwnerWith(user.IDEQ(owner)), server.NameEQ(name)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrServerNotFound
		}

		return nil, err
	}

	return srv, nil
}

type NodeResources struct {
	Memory       uint64
	Disk         uint64
	LogicalCores uint32
}

func GetAllocatedNodeResources(ctx context.Context, db *ent.Client, n *ent.Node) (*NodeResources, error) {
	allocatedRam, err := db.Server.Query().Where(server.HasNodeWith(node.IDEQ(n.ID))).Aggregate(ent.Sum(server.FieldRAM)).Int(ctx)
	if err != nil {
		return nil, err
	}

	allocatedStorage, err := db.Server.Query().Where(server.HasNodeWith(node.IDEQ(n.ID))).Aggregate(ent.Sum(server.FieldStorage)).Int(ctx)
	if err != nil {
		return nil, err
	}

	allocatedLogicalCores, err := db.Server.Query().Where(server.HasNodeWith(node.IDEQ(n.ID))).Aggregate(ent.Sum(server.FieldLogicalCores)).Int(ctx)
	if err != nil {
		return nil, err
	}

	res := NodeResources{
		Memory:       uint64(allocatedRam),
		Disk:         uint64(allocatedStorage),
		LogicalCores: uint32(allocatedLogicalCores),
	}

	return &res, nil
}

func GetFreeNodeResources(ctx context.Context, db *ent.Client, n *ent.Node) (*NodeResources, error) {
	//var allocatedRam int

	count, err := db.Server.Query().Where(server.HasNodeWith(node.IDEQ(n.ID))).Count(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &NodeResources{
			n.RAM - config.Config.Skyhook.MinFreeRAM,
			n.Storage - config.Config.Skyhook.MinFreeDisk,
			uint32(n.LogicalCores) - config.Config.Skyhook.MinFreeLogicalCores,
		}, nil
	}

	allocatedRam, err := db.Server.Query().Where(server.HasNodeWith(node.IDEQ(n.ID))).Aggregate(ent.Sum(server.FieldRAM)).Int(ctx)
	if err != nil {
		log.Printf("Failed getting allocated RAM: %v\n", err)
		return nil, err
	}

	allocatedStorage, err := db.Server.Query().Where(server.HasNodeWith(node.IDEQ(n.ID))).Aggregate(ent.Sum(server.FieldStorage)).Int(ctx)
	if err != nil {
		log.Printf("Failed getting allocated storage: %v\n", err)
		return nil, err
	}

	allocatedLogicalCores, err := db.Server.Query().Where(server.HasNodeWith(node.IDEQ(n.ID))).Aggregate(ent.Sum(server.FieldLogicalCores)).Int(ctx)
	if err != nil {
		log.Printf("Failed getting allocated cores: %v\n", err)
		return nil, err
	}

	res := NodeResources{
		Memory:       n.RAM - uint64(allocatedRam) - config.Config.Skyhook.MinFreeRAM,
		Disk:         n.Storage - uint64(allocatedStorage) - config.Config.Skyhook.MinFreeDisk,
		LogicalCores: uint32(n.LogicalCores) - uint32(allocatedLogicalCores) - config.Config.Skyhook.MinFreeLogicalCores,
	}

	return &res, nil
}

func HasEnoughFreeResources(ctx context.Context, db *ent.Client, n *ent.Node, requiredRes NodeResources) (bool, *NodeResources) {
	free, err := GetFreeNodeResources(ctx, db, n)
	if err != nil {
		log.Printf("Failed calculating node's free resources: %v", err)
		return false, nil
	}
	if free.Memory < requiredRes.Memory {
		return false, nil
	}
	if free.Disk < requiredRes.Disk {
		return false, nil
	}
	if free.LogicalCores < requiredRes.LogicalCores {
		return false, nil
	}

	return true, free
}

func NodeWithMostFreeResources(nodes map[*ent.Node]*NodeResources) *ent.Node {
	var minNode = struct {
		avg  float64
		node *ent.Node
	}{avg: math.Inf(1), node: nil}

	for k, v := range nodes {
		ramPercent := v.Memory / k.RAM
		storagePercent := v.Disk / k.Storage
		coresPercent := uint64(v.LogicalCores / uint32(k.LogicalCores))

		avg := (ramPercent + storagePercent + coresPercent) / 3
		if float64(avg) < minNode.avg {
			minNode.node = k
			minNode.avg = float64(avg)
		}
	}

	return minNode.node
}

func FindOptimalNode(ctx context.Context, db *ent.Client, opts *protoapi.ServersCreateRequest) (*ent.Node, error) {
	nodes, err := db.Node.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	requiredRes := NodeResources{
		Memory:       opts.Ram,
		Disk:         opts.Storage,
		LogicalCores: opts.LogicalCores,
	}

	freeNodes := make(map[*ent.Node]*NodeResources)
	for _, n := range nodes {
		if has, free := HasEnoughFreeResources(ctx, db, n, requiredRes); has {
			freeNodes[n] = free
		}
	}
	if len(freeNodes) == 0 {
		return nil, ErrNoFreeNodesFound
	}

	optimal := NodeWithMostFreeResources(freeNodes)
	if optimal == nil {
		return nil, ErrNoFreeNodesFound
	}

	return optimal, nil
}

func FindServerByID(ctx context.Context, db *ent.Client, serverId uuid.UUID) (*ent.Server, error) {
	srv, err := db.Server.Query().Where(server.IDEQ(serverId)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrServerNotFound
		}

		return nil, err
	}

	return srv, nil
}

func StartServer(ctx context.Context, store *module.Store, db *ent.Client, serverId uuid.UUID) error {
	srv, err := FindServerByID(ctx, db, serverId)
	fmt.Printf("%v\n", srv)
	if err != nil {
		return err
	}

	var client struct {
		StartServer func(srv protoapi.Server) error
	}

	variant := FindStoreCraterVariant(store, srv.Crater, srv.CraterVariant)
	if variant == nil {
		return ErrUnsupportedVariant
	}

	closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("http://localhost:%v", variant.Crater.Provider.Backend.RPCPort), "CratersHandler", &client, nil)
	if err != nil {
		return err
	}
	defer closer()

	err = client.StartServer(*proto.EntServerToProtoServer(*srv, store))
	if err != nil {
		return err
	}

	return nil
}

func StopServer(ctx context.Context, store *module.Store, db *ent.Client, serverId uuid.UUID) error {
	srv, err := FindServerByID(ctx, db, serverId)
	if err != nil {
		return err
	}

	var client struct {
		StopServer func(srv protoapi.Server) error
	}

	variant := FindStoreCraterVariant(store, srv.Crater, srv.CraterVariant)
	if variant == nil {
		return ErrUnsupportedVariant
	}

	closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("http://localhost:%v", variant.Crater.Provider.Backend.RPCPort), "CratersHandler", &client, nil)
	if err != nil {
		return err
	}
	defer closer()

	err = client.StopServer(*proto.EntServerToProtoServer(*srv, store))
	if err != nil {
		return err
	}

	return nil
}

func RestartServer(ctx context.Context, store *module.Store, db *ent.Client, serverId uuid.UUID) error {
	srv, err := FindServerByID(ctx, db, serverId)
	if err != nil {
		return err
	}

	var client struct {
		RestartServer func(srv protoapi.Server) error
	}

	variant := FindStoreCraterVariant(store, srv.Crater, srv.CraterVariant)
	if variant == nil {
		return ErrUnsupportedVariant
	}

	closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("http://localhost:%v", variant.Crater.Provider.Backend.RPCPort), "CratersHandler", &client, nil)
	if err != nil {
		return err
	}
	defer closer()

	err = client.RestartServer(*proto.EntServerToProtoServer(*srv, store))
	if err != nil {
		return err
	}

	return nil
}

func FindAllServers(ctx context.Context, db *ent.Client, _ *protoapi.ServersFindAllRequest) ([]*ent.Server, error) {
	srvs, err := db.Server.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return srvs, nil
}

func FindOneServer(ctx context.Context, db *ent.Client, req *protoapi.ServersFindOneRequest) (*ent.Server, error) {
	srv, err := db.Server.Query().Where(server.IDEQ(proto.ProtoUUIDToUUID(req.Id))).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrServerNotFound
		}

		return nil, err
	}

	return srv, nil
}

func CreateNodeDockerClient(ctx context.Context, db *ent.Client, serverId *protoapi.UUID) (*docker.Client, error) {
	srvId := proto.ProtoUUIDToUUID(serverId)

	srv, err := FindServerByID(ctx, db, srvId)
	if err != nil {
		return nil, err
	}
	nd, err := srv.QueryNode().First(ctx)
	if err != nil {
		return nil, err
	}

	host := fmt.Sprintf("tcp://%s:2375", nd.Ipv4Address)
	cli, err := docker.NewClientWithOpts(docker.WithHost(host), docker.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func DeleteOneServer(ctx context.Context, db *ent.Client, req *protoapi.ServersDeleteRequest) error {
	/*	srvId := proto.ProtoUUIDToUUID(req.Id)

		srv, err := FindServerByID(ctx, db, srvId)
		if err != nil {
			return err
		}
		nd, err := srv.QueryNode().First(ctx)
		if err != nil {
			return err
		}

		host := fmt.Sprintf("tcp://%s:2375", nd.Ipv4Address)
		cli, err := docker.NewClientWithOpts(docker.WithHost(host), docker.WithAPIVersionNegotiation())
		if err != nil {
			return err
		}*/
	cli, err := CreateNodeDockerClient(ctx, db, req.Id)
	if err != nil {
		return err
	}
	defer cli.Close()

	srvId := proto.ProtoUUIDToUUID(req.Id)

	srv, err := FindServerByID(ctx, db, srvId)
	if err != nil {
		return err
	}

	err = cli.ContainerRemove(ctx, srv.ContainerId, container.RemoveOptions{RemoveVolumes: true, Force: true})
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	err = db.Server.DeleteOneID(proto.ProtoUUIDToUUID(req.Id)).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrServerNotFound
		}

		return err
	}

	return nil
}

func InspectServerContainer(ctx context.Context, db *ent.Client, id *protoapi.UUID) (*types.ContainerJSON, error) {
	cli, err := CreateNodeDockerClient(ctx, db, id)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	srv, err := FindServerByID(ctx, db, proto.ProtoUUIDToUUID(id))
	if err != nil {
		return nil, err
	}

	info, err := cli.ContainerInspect(ctx, srv.ContainerId)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
