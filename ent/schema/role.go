package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

// Role holds the schema definition for the Role entity.
type Role struct {
    ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now),
        field.Time("deleted_at").Optional(),
        field.String("name").MaxLen(24).Unique(),
        field.Strings("permissions").Optional(),
    }
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
    return nil
}
