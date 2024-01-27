// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/jswildcards/gocareer/schema"
)

func main() {
	r := gin.Default()

	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})

	r.POST("/graphql", gin.WrapH(h))

	r.Run(":3000")
}
