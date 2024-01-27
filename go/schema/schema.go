package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/jswildcards/gocareer/schema/query"
)

var schemaConfig = graphql.SchemaConfig{
	Query: query.QueryObject,
}

var Schema, _ = graphql.NewSchema(schemaConfig)
