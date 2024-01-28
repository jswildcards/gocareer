package login_controller

import (
	"github.com/graphql-go/graphql"
	"github.com/jswildcards/gocareer/service/login_service"
)

func Login(p graphql.ResolveParams) (interface{}, error) {
	return login_service.Login(), nil
}
