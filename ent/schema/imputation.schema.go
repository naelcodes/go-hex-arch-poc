package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Imputation struct {
	ent.Schema
}

func (Imputation) Fields() []ent.Field {

	return []ent.Field{

		field.Float("amount_apply").
			ValueScanner(Money{CurrencyPrefix: "$"}).
			SchemaType(map[string]string{
				dialect.Postgres: "money",
			}).Min(0).Default(0),
		field.Enum("tag").Values("1", "2", "3").Default("3"),
	}
}

func (Imputation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("invoice", Invoice.Type).
			Ref("imputations").Unique().
			Required(),
		edge.From("payment", Payment.Type).
			Ref("imputations").Unique().Required(),
	}
}

func (Imputation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invoice_payment_received"},
	}
}
