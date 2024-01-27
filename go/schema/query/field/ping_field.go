package ping_field

import (
	"github.com/graphql-go/graphql"
	ping_controller "github.com/jswildcards/gocareer/controller"
)

var Field = &graphql.Field{
	Type:        graphql.String,
	Description: "ping server",
	Args:        nil,
	Resolve:     ping_controller.Ping,
}
