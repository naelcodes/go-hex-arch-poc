package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TravelItem struct {
	ent.Schema
}

// Fields of the TravelItem.
func (TravelItem) Fields() []ent.Field {

	return []ent.Field{

		field.Float("total_price").SchemaType(map[string]string{
			dialect.Postgres: "money",
		}).Positive().Default(0),
		field.String("itinerary"),
		field.String("traveler_name"),
		field.String("ticket_number"),
		field.Int("conjunction_number"),
		field.Enum("status").Values("pending", "invoiced", "void", "receipted").Default("pending"),
	}
}

// Edges of the TravelItem.
func (TravelItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("invoice", Invoice.Type).
			Ref("travel_items").Unique().
			Required(),
	}
}

// Annotations of the TravelItem.
func (TravelItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "air_booking"},
	}
}
