package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Customer struct {
	ent.Schema
}

func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_name"),
		field.String("account_number").NotEmpty().Unique(),
		field.Int("id_currency").Default(550),
		field.Int("id_country").Default(40),
		field.String("alias").NotEmpty().Unique(),
		field.String("ab_key").NotEmpty().Unique(),
		field.String("state").Optional(),
		field.String("tmc_client_number").NotEmpty().Unique(),
		field.Enum("tag").Values("1", "2", "3").Default("3"),
	}
}

func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invoices", Invoice.Type).
			StorageKey(edge.Column("id_customer")),

		edge.To("payments", Payment.Type).
			StorageKey(edge.Column("id_customer")),
	}
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "customer"},
	}
}
