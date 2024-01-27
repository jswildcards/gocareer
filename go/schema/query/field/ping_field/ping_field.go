package ping_field

import (
	"github.com/graphql-go/graphql"
	"github.com/jswildcards/gocareer/controller/ping_controller"
)

var Field = &graphql.Field{
	Type:        graphql.String,
	Description: "ping server",
	Args:        nil,
	Resolve:     ping_controller.Ping,
}
