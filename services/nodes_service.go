package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/ent/node"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
	"github.com/Encedeus/panel/validate"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"strconv"
	"time"
)

const SKYHOOK_PORT = 8000

type OS string

const (
	UBUNTU_22_04_LTS OS = "ubuntu-22.04-lts"
	UBUNTU_24_04_LTS OS = "ubuntu-24.04-lts"
)

type SSHCredentials struct {
	Username   string
	Password   string
	PublicKey  string
	PrivateKey string
	Passphrase string
	Port       module.Port
}

func CreateNode(ctx context.Context, db *ent.Client, req *protoapi.NodesCreateRequest) (*protoapi.NodesCreateResponse, error) {
	n, err := FindNodeByIP(ctx, db, req.Ipv4Address)
	if errors.Is(err, ErrInvalidIPAddress) {
		return nil, err
	}
	if n != nil {
		return nil, ErrNodeAlreadyExists
	}

	n, err = FindNodeByFQDN(ctx, db, req.Fqdn, req.Ipv4Address)
	/*	if errors.Is(err, ErrInvalidFqdn) {
		return nil, err
	}*/
	if n != nil {
		return nil, ErrNodeAlreadyExists
	}
	//if !validate.IsDomain(req.Fqdn, req.Ipv4Address) {
	//	return nil, ErrInvalidFqdn
	//}

	// TODO: add automatic Skyhook installation

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(timeoutCtx, fmt.Sprintf("%v:%v", req.Ipv4Address, SKYHOOK_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, ErrFailedConnectingToSkyhook
	}
	defer conn.Close()

	timeoutCtx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	hp := protoapi.NewNodeInfoClient(conn)
	hwInfo, err := hp.GetNodeHardwareInfo(timeoutCtx, &protoapi.HardwareInfoRequest{})
	if err != nil {
		return nil, ErrFailedGettingHardwareInfo
	}
	fmt.Printf("Hardware: %+v\n", hwInfo)

	n, err = db.Node.Create().
		SetIpv4Address(req.Ipv4Address).
		SetFqdn(req.Fqdn).
		//SetSkyhookVersion(req).
		SetOs(hwInfo.Os).
		SetCPU(hwInfo.Cpu).
		SetCPUBaseClock(uint(hwInfo.CpuClockSpeed)).
		SetCores(uint(hwInfo.Cores)).
		SetLogicalCores(uint(hwInfo.LogicalCores)).
		SetRAM(hwInfo.TotalMemory).
		SetStorage(hwInfo.TotalDisk).
		SetSkyhookAPIKey(req.ApiKey).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	resp := &protoapi.NodesCreateResponse{
		Nodes: []*protoapi.Node{
			{
				Id:             proto.UUIDToProtoUUID(n.ID),
				CreatedAt:      timestamppb.New(n.CreatedAt),
				UpdatedAt:      timestamppb.New(n.UpdatedAt),
				Ipv4Address:    n.Ipv4Address,
				Fqdn:           n.Fqdn,
				SkyhookVersion: n.SkyhookVersion,
				Os:             n.Os,
				Cpu:            n.CPU,
				CpuBaseClock:   uint32(n.CPUBaseClock),
				Cores:          uint32(n.Cores),
				LogicalCores:   uint32(n.LogicalCores),
				Ram:            strconv.FormatUint(n.RAM, 10),
				Storage:        strconv.FormatUint(n.Storage, 10),
			},
		},
	}
	fmt.Printf("Response: %+v\n", resp)

	return resp, nil
}

func FindNodeByIP(ctx context.Context, db *ent.Client, ipAddress string) (*ent.Node, error) {
	if !validate.IsIpv4Address(ipAddress) {
		return nil, ErrInvalidIPAddress
	}

	n, err := db.Node.Query().Where(node.Ipv4AddressEQ(ipAddress)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrNodeNotFound
	}
	if err != nil {
		return nil, err
	}

	return n, nil
}

func FindNodeByFQDN(ctx context.Context, db *ent.Client, fqdn, ipAddress string) (*ent.Node, error) {
	if !validate.IsDomain(fqdn, ipAddress) {
		return nil, ErrInvalidFqdn
	}

	n, err := db.Node.Query().Where(node.FqdnEQ(fqdn)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrNodeNotFound
	}
	if err != nil {
		return nil, err
	}

	return n, nil
}

func FindNodeByID(ctx context.Context, db *ent.Client, req *protoapi.NodesFindOneRequest) (*protoapi.NodesFindOneResponse, error) {
	n, err := db.Node.Query().Where(node.IDEQ(proto.ProtoUUIDToUUID(req.Id))).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrNodeNotFound
		}
		return nil, err
	}

	resp := &protoapi.NodesFindOneResponse{
		Nodes: []*protoapi.Node{
			proto.EntNodeToProtoNode(*n),
		},
	}

	return resp, nil
}

func FindAllNodes(ctx context.Context, db *ent.Client, req *protoapi.NodesFindAllRequest) (*protoapi.NodesFindAllResponse, error) {
	dbNodes, err := db.Node.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	nodes := make([]*protoapi.Node, len(dbNodes))
	for idx, dn := range dbNodes {
		n := proto.EntNodeToProtoNode(*dn)
		nodes[idx] = n
	}

	resp := &protoapi.NodesFindAllResponse{
		Nodes: nodes,
	}

	return resp, nil
}

func DeleteNode(ctx context.Context, db *ent.Client, req *protoapi.NodesDeleteRequest) (*protoapi.NodesDeleteResponse, error) {
	err := db.Node.DeleteOneID(proto.ProtoUUIDToUUID(req.Id)).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrNodeNotFound
		}
		return nil, err
	}

	resp := &protoapi.NodesDeleteResponse{}

	return resp, nil
}

func InstallSkyhook(creds SSHCredentials, host string, port module.Port) error {
	hostKey, err := ssh.ParsePublicKey([]byte(creds.PublicKey))
	if err != nil {
		return ErrInvalidPublicKey
	}

	signer, err := ssh.ParsePrivateKeyWithPassphrase([]byte(creds.PrivateKey), []byte(creds.Passphrase))
	if err != nil {
		return ErrInvalidPrivateKeyOrPassphrase
	}

	config := &ssh.ClientConfig{
		User: creds.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(creds.Password),
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(host, strconv.Itoa(int(port))), config)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedConnectingToSSH, err)
	}
	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedCreatingSSHSession, err)
	}
	defer sess.Close()
	//sess.

	return nil
}
