package schema

import (
	"github.com/graphql-go/graphql"
	root_query "github.com/jswildcards/gocareer/schema/query"
)

var schemaConfig = graphql.SchemaConfig{
	Query: root_query.QueryObject,
}

var Schema, _ = graphql.NewSchema(schemaConfig)
