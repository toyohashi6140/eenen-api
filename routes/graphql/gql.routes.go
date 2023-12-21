package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/toyohashi6140/eenen-api/adapters/resolvers"
	"github.com/toyohashi6140/eenen-api/infra/graphql"
)

func DefineGraphQL(r *resolvers.Resolver) echo.HandlerFunc {
	// gqlgenで生成されたgraphqlハンドラを作成
	config := graphql.Config{
		Resolvers: r,
	}
	handler := handler.NewDefaultServer(graphql.NewExecutableSchema(config))
	return func(c echo.Context) error {
		handler.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
