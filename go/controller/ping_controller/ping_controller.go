package ping_controller

import (
	"github.com/graphql-go/graphql"
	"github.com/jswildcards/gocareer/service/ping_service"
)

func Ping(p graphql.ResolveParams) (interface{}, error) {
	return ping_service.Ping(), nil
}
