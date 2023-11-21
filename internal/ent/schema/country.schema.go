package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Country struct {
	ent.Schema
}

func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
	}
}

func (Country) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "country"},
	}
}
