package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

// ApiKey holds the schema definition for the ApiKey entity.
type ApiKey struct {
    ent.Schema
}

// Fields of the ApiKey.
func (ApiKey) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
        field.String("description").Optional(),
        field.Strings("ip_addresses").Optional(),
        field.String("key").NotEmpty(),
        field.UUID("user_id", uuid.UUID{}),
    }
}

// Edges of the ApiKey.
func (ApiKey) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Field("user_id").Required().Unique(),
    }
}
