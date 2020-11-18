package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Object holds the schema definition for the Object entity.
type Object struct {
	ent.Schema
}

// Fields of the Object.
func (Object) Fields() []ent.Field {
	return []ent.Field{
		// Collisions are okay here, as we search in pair of owner:object
		field.UUID("object_id", uuid.UUID{}).
			Default(uuid.New),
		field.Bytes("data"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Nillable().
			Optional(),
	}
}

// Edges of the Object.
func (Object) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("objects").
			Unique(),
		edge.To("past_objects", PastObject.Type),
	}
}
