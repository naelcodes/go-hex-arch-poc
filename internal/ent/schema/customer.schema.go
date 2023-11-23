package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Customer struct {
	ent.Schema
}

func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_name"),
		field.String("account_number").Unique(),
		field.Int("id_currency").Default(550),
		field.Int("id_country").Default(40),
		field.String("alias").Unique(),
		field.String("ab_key").Unique(),
		field.String("state").Optional(),
		field.String("tmc_client_number").Unique(),
		field.Enum("Tag").Values("1", "2", "3").Default("3"),
	}
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "customer"},
	}
}
