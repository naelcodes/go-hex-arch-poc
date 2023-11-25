package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Invoice struct {
	ent.Schema
}

func (Invoice) Fields() []ent.Field {

	return []ent.Field{

		field.String("creation_date").SchemaType(map[string]string{
			dialect.Postgres: "timestamp without timezone",
		}).NotEmpty(),
		field.String("invoice_number").NotEmpty().Unique(),
		field.Enum("status").Values("draft", "paid", "overdue", "unpaid", "void").
			Default("unpaid"),
		field.String("due_date").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Float("amount").SchemaType(map[string]string{
			dialect.Postgres: "money",
		}).Positive().Default(0),
		field.Float("balance").SchemaType(map[string]string{
			dialect.Postgres: "money",
		}).Positive().Default(0),
		field.Float("credit_apply").SchemaType(map[string]string{
			dialect.Postgres: "money",
		}).Positive().Default(0),
		field.Enum("tag").Values("1", "2", "3").Default("3"),
	}
}

func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("invoices").Unique().
			Required(),
		edge.To("imputations", Imputation.Type).
			StorageKey(edge.Column("id_invoice")),
		edge.To("travel_items", TravelItem.Type).
			StorageKey(edge.Column("id_invoice")),
	}
}

func (Invoice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invoice"},
	}
}
