// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {
	r := gin.Default()

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	r.POST("/graphql", gin.WrapH(h))

	r.Run(":3000")
}
