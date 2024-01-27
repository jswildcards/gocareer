package root_query

import (
	"github.com/graphql-go/graphql"
	ping_field "github.com/jswildcards/gocareer/schema/query/field"
)

var fields = graphql.Fields{
	"ping": ping_field.Field,
}

var QueryObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	},
)
