package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH"},
		AllowCredentials: true,
	}))
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"graphql-playground": "http://localhost:8080/graphiql",
			"graphql":            "http://localhost:8080/graphql",
			"v1":                 "http://localhost:8080/api/v1",
		})
	})

	V1(r)
	Graphql(r)

	return r
}

func Run() *gin.Engine {
	r := SetupRouter()
	err := r.Run()
	if err != nil {
		return nil
	}

	return r
}
