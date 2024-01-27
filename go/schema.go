package main

import (
	"github.com/graphql-go/graphql"
)

var fields = graphql.Fields{
	"ping": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "ok", nil
		},
	},
}

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	},
)

var schemaConfig = graphql.SchemaConfig{
	Query: rootQuery,
}

var schema, _ = graphql.NewSchema(schemaConfig)
