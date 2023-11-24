package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Payment struct {
	ent.Schema
}

func (Payment) Fields() []ent.Field {
	return []ent.Field{

		field.String("number").NotEmpty().Unique(),
		field.String("date").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).NotEmpty(),
		field.Float("balance").SchemaType(map[string]string{
			dialect.Postgres: "money",
		}).Positive().Default(0),
		field.Float("amount").SchemaType(map[string]string{
			dialect.Postgres: "money",
		}).Positive().Default(0),
		field.Enum("fop").Values("cash", "check", "bank_transfer").Default("cash"),
		field.Float("used_amount").SchemaType(map[string]string{
			dialect.Postgres: "money",
		}).Positive().Default(0),
		field.Enum("status").Values("open", "used", "void").Default("open"),
		field.Int("id_charts_of_accounts").Default(39),
		field.Int("id_currency").Default(550),
		field.Enum("Tag").Values("1", "2", "3").Default("3"),

		field.Int("id_payment_received").Optional(),
	}

}

func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("payments").Unique().
			Required(),
		edge.To("imputations", Imputation.Type).
			StorageKey(edge.Column("id_payment_received")),
	}
}

func (Payment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "payment_received"},
	}
}
