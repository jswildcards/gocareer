package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/jswildcards/gocareer/schema/mutation"
	"github.com/jswildcards/gocareer/schema/query"
)

var schemaConfig = graphql.SchemaConfig{
	Query:    query.QueryObject,
	Mutation: mutation.MutationObject,
}

var Schema, _ = graphql.NewSchema(schemaConfig)
