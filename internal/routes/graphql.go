package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-graphql/graph"
	"github.com/onosproject/aether-graphql/graph/generated"
)

func Graphql(router *gin.Engine) *gin.Engine {
	router.POST("/graphql", graphqlHandler())
	router.GET("/graphiql", playgroundHandler())

	return router
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {

	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
