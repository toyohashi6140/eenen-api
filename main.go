package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/toyohashi6140/eenen-api/adapters/resolvers"
	gqlroutes "github.com/toyohashi6140/eenen-api/routes/graphql"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.POST("/query", gqlroutes.DefineGraphQL(resolvers.NewResolver()))

	log.Fatal(e.Start(":1323"))
}
