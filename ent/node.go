// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Encedeus/panel/ent/node"
	"github.com/google/uuid"
)

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Ipv4Address holds the value of the "ipv4_address" field.
	Ipv4Address string `json:"ipv4_address,omitempty"`
	// Fqdn holds the value of the "fqdn" field.
	Fqdn string `json:"fqdn,omitempty"`
	// SkyhookVersion holds the value of the "skyhook_version" field.
	SkyhookVersion string `json:"skyhook_version,omitempty"`
	// SkyhookAPIKey holds the value of the "skyhook_api_key" field.
	SkyhookAPIKey string `json:"skyhook_api_key,omitempty"`
	// Os holds the value of the "os" field.
	Os string `json:"os,omitempty"`
	// CPU holds the value of the "cpu" field.
	CPU string `json:"cpu,omitempty"`
	// CPUBaseClock holds the value of the "cpu_base_clock" field.
	CPUBaseClock uint `json:"cpu_base_clock,omitempty"`
	// Cores holds the value of the "cores" field.
	Cores uint `json:"cores,omitempty"`
	// LogicalCores holds the value of the "logical_cores" field.
	LogicalCores uint `json:"logical_cores,omitempty"`
	// RAM holds the value of the "ram" field.
	RAM uint64 `json:"ram,omitempty"`
	// Storage holds the value of the "storage" field.
	Storage uint64 `json:"storage,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeQuery when eager-loading is set.
	Edges        NodeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NodeEdges holds the relations/edges for other nodes in the graph.
type NodeEdges struct {
	// Nodes holds the value of the nodes edge.
	Nodes []*Server `json:"nodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) NodesOrErr() ([]*Server, error) {
	if e.loadedTypes[0] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Node) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case node.FieldCPUBaseClock, node.FieldCores, node.FieldLogicalCores, node.FieldRAM, node.FieldStorage:
			values[i] = new(sql.NullInt64)
		case node.FieldIpv4Address, node.FieldFqdn, node.FieldSkyhookVersion, node.FieldSkyhookAPIKey, node.FieldOs, node.FieldCPU:
			values[i] = new(sql.NullString)
		case node.FieldCreatedAt, node.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case node.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Node fields.
func (n *Node) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				n.ID = *value
			}
		case node.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case node.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case node.FieldIpv4Address:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ipv4_address", values[i])
			} else if value.Valid {
				n.Ipv4Address = value.String
			}
		case node.FieldFqdn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fqdn", values[i])
			} else if value.Valid {
				n.Fqdn = value.String
			}
		case node.FieldSkyhookVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field skyhook_version", values[i])
			} else if value.Valid {
				n.SkyhookVersion = value.String
			}
		case node.FieldSkyhookAPIKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field skyhook_api_key", values[i])
			} else if value.Valid {
				n.SkyhookAPIKey = value.String
			}
		case node.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				n.Os = value.String
			}
		case node.FieldCPU:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cpu", values[i])
			} else if value.Valid {
				n.CPU = value.String
			}
		case node.FieldCPUBaseClock:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu_base_clock", values[i])
			} else if value.Valid {
				n.CPUBaseClock = uint(value.Int64)
			}
		case node.FieldCores:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cores", values[i])
			} else if value.Valid {
				n.Cores = uint(value.Int64)
			}
		case node.FieldLogicalCores:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field logical_cores", values[i])
			} else if value.Valid {
				n.LogicalCores = uint(value.Int64)
			}
		case node.FieldRAM:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ram", values[i])
			} else if value.Valid {
				n.RAM = uint64(value.Int64)
			}
		case node.FieldStorage:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field storage", values[i])
			} else if value.Valid {
				n.Storage = uint64(value.Int64)
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Node.
// This includes values selected through modifiers, order, etc.
func (n *Node) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryNodes queries the "nodes" edge of the Node entity.
func (n *Node) QueryNodes() *ServerQuery {
	return NewNodeClient(n.config).QueryNodes(n)
}

// Update returns a builder for updating this Node.
// Note that you need to call Node.Unwrap() before calling this method if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return NewNodeClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Node entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Node) Unwrap() *Node {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Node is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString("Node(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ipv4_address=")
	builder.WriteString(n.Ipv4Address)
	builder.WriteString(", ")
	builder.WriteString("fqdn=")
	builder.WriteString(n.Fqdn)
	builder.WriteString(", ")
	builder.WriteString("skyhook_version=")
	builder.WriteString(n.SkyhookVersion)
	builder.WriteString(", ")
	builder.WriteString("skyhook_api_key=")
	builder.WriteString(n.SkyhookAPIKey)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(n.Os)
	builder.WriteString(", ")
	builder.WriteString("cpu=")
	builder.WriteString(n.CPU)
	builder.WriteString(", ")
	builder.WriteString("cpu_base_clock=")
	builder.WriteString(fmt.Sprintf("%v", n.CPUBaseClock))
	builder.WriteString(", ")
	builder.WriteString("cores=")
	builder.WriteString(fmt.Sprintf("%v", n.Cores))
	builder.WriteString(", ")
	builder.WriteString("logical_cores=")
	builder.WriteString(fmt.Sprintf("%v", n.LogicalCores))
	builder.WriteString(", ")
	builder.WriteString("ram=")
	builder.WriteString(fmt.Sprintf("%v", n.RAM))
	builder.WriteString(", ")
	builder.WriteString("storage=")
	builder.WriteString(fmt.Sprintf("%v", n.Storage))
	builder.WriteByte(')')
	return builder.String()
}

// Nodes is a parsable slice of Node.
type Nodes []*Node
