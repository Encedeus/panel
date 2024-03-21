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
		field.String("skyhook_version").Optional(),
		field.String("skyhook_api_key").Unique(),
		field.String("os").Optional(),
		field.String("cpu").Optional(),
		// in MHz
		field.Uint("cpu_base_clock").Optional(),
		field.Uint("cores").Optional(),
		field.Uint("logical_cores").Optional(),
		// In KB
		field.Uint64("ram").Optional(),
		// In KB
		field.Uint64("storage").Optional(),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nodes", Server.Type),
	}
}
