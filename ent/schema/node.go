package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now),
		field.String("ipv4_address").Unique(),
		field.String("fqdn").Unique(),
		field.String("skyhook_version"),
		field.String("os"),
		field.String("cpu"),
		// in MHz
		field.Uint("cpu_base_clock"),
		field.Uint("cores"),
		field.Uint("logical_cores"),
		// In KB
		field.Uint("ram"),
		// In KB
		field.Uint("storage"),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nodes", Server.Type),
	}
}
