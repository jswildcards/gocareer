package login_field

import (
	"github.com/graphql-go/graphql"
	"github.com/jswildcards/gocareer/controller/login_controller"
)

var Field = &graphql.Field{
	Type:        graphql.String,
	Description: "login server",
	Args:        nil,
	Resolve:     login_controller.Login,
}
