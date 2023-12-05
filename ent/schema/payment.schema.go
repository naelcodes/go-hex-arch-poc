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
		field.Float("balance").
			ValueScanner(Money{CurrencyPrefix: "$"}).
			SchemaType(map[string]string{
				dialect.Postgres: "money",
			}).Min(0).Default(0),
		field.Float("amount").
			ValueScanner(Money{CurrencyPrefix: "$"}).
			SchemaType(map[string]string{
				dialect.Postgres: "money",
			}).Min(0).Default(0),
		field.Float("base_amount").
			ValueScanner(Money{CurrencyPrefix: "$"}).
			SchemaType(map[string]string{
				dialect.Postgres: "money",
			}).Min(0).Default(0),
		field.Float("used_amount").
			ValueScanner(Money{CurrencyPrefix: "$"}).
			SchemaType(map[string]string{
				dialect.Postgres: "money",
			}).Min(0).Default(0),

		field.Enum("type").Values("supplier_refund", "transfer_from_account", "other_income", "customer_payment", "sales_receipt").Default("customer_payment"),
		field.Enum("fop").Values("cash", "check", "bank_transfer").Default("cash"),

		field.Enum("status").Values("open", "used", "void").Default("open"),
		field.Int("id_chart_of_accounts").Default(39),
		field.Int("id_currency").Default(550),
		field.Enum("Tag").Values("1", "2", "3").Default("3"),
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
