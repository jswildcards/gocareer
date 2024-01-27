package ping_controller

import (
	"github.com/graphql-go/graphql"

	ping_service "github.com/jswildcards/gocareer/service"
)

func Ping(p graphql.ResolveParams) (interface{}, error) {
	return ping_service.Ping(), nil
}
