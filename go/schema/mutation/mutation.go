package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/jswildcards/gocareer/schema/mutation/field/login_field"
)

var fields = graphql.Fields{
	"login": login_field.Field,
}

var MutationObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: fields,
	},
)
