package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// PastObject holds the schema definition for the PastObject entity.
type PastObject struct {
	ent.Schema
}

// Fields of the PastObject.
func (PastObject) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("data"),
		field.Time("replaced_at").
			Default(time.Now),
	}
}

// Edges of the PastObject.
func (PastObject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("present_object", Object.Type).
			Ref("past_objects").
			Unique(),
	}
}
