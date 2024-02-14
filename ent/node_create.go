// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/panel/ent/node"
	"github.com/Encedeus/panel/ent/server"
	"github.com/google/uuid"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	mutation *NodeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (nc *NodeCreate) SetCreatedAt(t time.Time) *NodeCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NodeCreate) SetNillableCreatedAt(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NodeCreate) SetUpdatedAt(t time.Time) *NodeCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NodeCreate) SetNillableUpdatedAt(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetIpv4Address sets the "ipv4_address" field.
func (nc *NodeCreate) SetIpv4Address(s string) *NodeCreate {
	nc.mutation.SetIpv4Address(s)
	return nc
}

// SetFqdn sets the "fqdn" field.
func (nc *NodeCreate) SetFqdn(s string) *NodeCreate {
	nc.mutation.SetFqdn(s)
	return nc
}

// SetSkyhookVersion sets the "skyhook_version" field.
func (nc *NodeCreate) SetSkyhookVersion(s string) *NodeCreate {
	nc.mutation.SetSkyhookVersion(s)
	return nc
}

// SetOs sets the "os" field.
func (nc *NodeCreate) SetOs(s string) *NodeCreate {
	nc.mutation.SetOs(s)
	return nc
}

// SetCPU sets the "cpu" field.
func (nc *NodeCreate) SetCPU(s string) *NodeCreate {
	nc.mutation.SetCPU(s)
	return nc
}

// SetCPUBaseClock sets the "cpu_base_clock" field.
func (nc *NodeCreate) SetCPUBaseClock(u uint) *NodeCreate {
	nc.mutation.SetCPUBaseClock(u)
	return nc
}

// SetCores sets the "cores" field.
func (nc *NodeCreate) SetCores(u uint) *NodeCreate {
	nc.mutation.SetCores(u)
	return nc
}

// SetLogicalCores sets the "logical_cores" field.
func (nc *NodeCreate) SetLogicalCores(u uint) *NodeCreate {
	nc.mutation.SetLogicalCores(u)
	return nc
}

// SetRAM sets the "ram" field.
func (nc *NodeCreate) SetRAM(u uint) *NodeCreate {
	nc.mutation.SetRAM(u)
	return nc
}

// SetStorage sets the "storage" field.
func (nc *NodeCreate) SetStorage(u uint) *NodeCreate {
	nc.mutation.SetStorage(u)
	return nc
}

// SetID sets the "id" field.
func (nc *NodeCreate) SetID(u uuid.UUID) *NodeCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NodeCreate) SetNillableID(u *uuid.UUID) *NodeCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// AddNodeIDs adds the "nodes" edge to the Server entity by IDs.
func (nc *NodeCreate) AddNodeIDs(ids ...uuid.UUID) *NodeCreate {
	nc.mutation.AddNodeIDs(ids...)
	return nc
}

// AddNodes adds the "nodes" edges to the Server entity.
func (nc *NodeCreate) AddNodes(s ...*Server) *NodeCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nc.AddNodeIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nc *NodeCreate) Mutation() *NodeMutation {
	return nc.mutation
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NodeCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NodeCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NodeCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := node.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := node.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := node.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NodeCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Node.created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Node.updated_at"`)}
	}
	if _, ok := nc.mutation.Ipv4Address(); !ok {
		return &ValidationError{Name: "ipv4_address", err: errors.New(`ent: missing required field "Node.ipv4_address"`)}
	}
	if _, ok := nc.mutation.Fqdn(); !ok {
		return &ValidationError{Name: "fqdn", err: errors.New(`ent: missing required field "Node.fqdn"`)}
	}
	if _, ok := nc.mutation.SkyhookVersion(); !ok {
		return &ValidationError{Name: "skyhook_version", err: errors.New(`ent: missing required field "Node.skyhook_version"`)}
	}
	if _, ok := nc.mutation.Os(); !ok {
		return &ValidationError{Name: "os", err: errors.New(`ent: missing required field "Node.os"`)}
	}
	if _, ok := nc.mutation.CPU(); !ok {
		return &ValidationError{Name: "cpu", err: errors.New(`ent: missing required field "Node.cpu"`)}
	}
	if _, ok := nc.mutation.CPUBaseClock(); !ok {
		return &ValidationError{Name: "cpu_base_clock", err: errors.New(`ent: missing required field "Node.cpu_base_clock"`)}
	}
	if _, ok := nc.mutation.Cores(); !ok {
		return &ValidationError{Name: "cores", err: errors.New(`ent: missing required field "Node.cores"`)}
	}
	if _, ok := nc.mutation.LogicalCores(); !ok {
		return &ValidationError{Name: "logical_cores", err: errors.New(`ent: missing required field "Node.logical_cores"`)}
	}
	if _, ok := nc.mutation.RAM(); !ok {
		return &ValidationError{Name: "ram", err: errors.New(`ent: missing required field "Node.ram"`)}
	}
	if _, ok := nc.mutation.Storage(); !ok {
		return &ValidationError{Name: "storage", err: errors.New(`ent: missing required field "Node.storage"`)}
	}
	return nil
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NodeCreate) createSpec() (*Node, *sqlgraph.CreateSpec) {
	var (
		_node = &Node{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(node.Table, sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID))
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(node.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(node.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.Ipv4Address(); ok {
		_spec.SetField(node.FieldIpv4Address, field.TypeString, value)
		_node.Ipv4Address = value
	}
	if value, ok := nc.mutation.Fqdn(); ok {
		_spec.SetField(node.FieldFqdn, field.TypeString, value)
		_node.Fqdn = value
	}
	if value, ok := nc.mutation.SkyhookVersion(); ok {
		_spec.SetField(node.FieldSkyhookVersion, field.TypeString, value)
		_node.SkyhookVersion = value
	}
	if value, ok := nc.mutation.Os(); ok {
		_spec.SetField(node.FieldOs, field.TypeString, value)
		_node.Os = value
	}
	if value, ok := nc.mutation.CPU(); ok {
		_spec.SetField(node.FieldCPU, field.TypeString, value)
		_node.CPU = value
	}
	if value, ok := nc.mutation.CPUBaseClock(); ok {
		_spec.SetField(node.FieldCPUBaseClock, field.TypeUint, value)
		_node.CPUBaseClock = value
	}
	if value, ok := nc.mutation.Cores(); ok {
		_spec.SetField(node.FieldCores, field.TypeUint, value)
		_node.Cores = value
	}
	if value, ok := nc.mutation.LogicalCores(); ok {
		_spec.SetField(node.FieldLogicalCores, field.TypeUint, value)
		_node.LogicalCores = value
	}
	if value, ok := nc.mutation.RAM(); ok {
		_spec.SetField(node.FieldRAM, field.TypeUint, value)
		_node.RAM = value
	}
	if value, ok := nc.mutation.Storage(); ok {
		_spec.SetField(node.FieldStorage, field.TypeUint, value)
		_node.Storage = value
	}
	if nodes := nc.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.NodesTable,
			Columns: []string{node.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	builders []*NodeCreate
}

// Save creates the Node entities in the database.
func (ncb *NodeCreateBulk) Save(ctx context.Context) ([]*Node, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Node, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NodeCreateBulk) SaveX(ctx context.Context) []*Node {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NodeCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}