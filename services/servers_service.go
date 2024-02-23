package services

import (
	"context"
	"fmt"
	"github.com/Encedeus/panel/config"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/ent/node"
	"github.com/Encedeus/panel/ent/server"
	"github.com/Encedeus/panel/ent/user"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/google/uuid"
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

	_, err = db.Node.Query().Where(node.IDEQ(proto.ProtoUUIDToUUID(req.Node))).First(ctx)
	if err != nil {
		return nil, ErrNodeNotFound
	}

	n, err := FindOptimalNode(ctx, db, req)
	if err != nil {
		return nil, err
	}
	req.Node = proto.UUIDToProtoUUID(n.ID)

	var client struct {
		CreateServer func(req *protoapi.ServersCreateRequest) (*protoapi.ServersCreateResponse, error)
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

	resp, err := client.CreateServer(req)
	if err != nil {
		return nil, err
	}

	_, err = db.Server.Create().
		SetName(req.Name).
		SetCraterProvider(variant.Crater.Provider.Manifest.Name).
		SetCrater(req.Crater).
		SetCraterVariant(req.CraterVariant).
		SetOwnerID(proto.ProtoUUIDToUUID(req.Owner)).
		SetRAM(req.Ram).
		SetStorage(req.Storage).
		SetLogicalCores(uint(req.LogicalCores)).
		SetPort(uint16(req.Port.Value)).
		SetNode(n).
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

	count, err := db.Server.Query().Count(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &NodeResources{
			n.RAM,
			n.Storage,
			uint32(n.LogicalCores),
		}, nil
	}

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
