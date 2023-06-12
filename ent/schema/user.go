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
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.String("email").MaxLen(32),
		field.String("password"),
		field.Bytes("pfp").MaxLen(4000000),
		field.String("name").MaxLen(32),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type),
	}
}