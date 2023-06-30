package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now),
		field.Time("deleted_at").Optional(),
		field.String("email").MaxLen(32),
		field.String("password"),
		field.String("name").MaxLen(32).Unique(),
		field.Int("role_id"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).Field("role_id").Unique().Required(),
	}
}
