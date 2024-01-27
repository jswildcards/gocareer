package query

import (
	"github.com/graphql-go/graphql"
	"github.com/jswildcards/gocareer/schema/query/field/ping_field"
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
