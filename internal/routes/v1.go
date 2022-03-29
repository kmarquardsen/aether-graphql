package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-graphql/internal/controllers"
)

func V1(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	enterprises(v1)

	return r
}

func enterprises(v1 *gin.RouterGroup) {
	{
		enterprise := new(controllers.Enterprise)
		v1.GET("/enterprises", enterprise.List)
		v1.POST("/enterprises", enterprise.Store)
		v1.PATCH("/enterprises", enterprise.Update)
		v1.DELETE("/enterprises/:id", enterprise.Delete)
		v1.GET("/enterprises/:id", enterprise.Get)
	}
}

func sites(v1 *gin.RouterGroup) {
	{
		site := new(controllers.Site)
		v1.GET("/sites", site.List)
		v1.POST("/sites", site.Store)
		v1.PATCH("/sites", site.Update)
		v1.DELETE("/sites/:id", site.Delete)
		v1.GET("/sites/:id", site.Get)
	}
}

func devices(v1 *gin.RouterGroup) {
	{
		device := new(controllers.Device)
		v1.GET("/devices", device.List)
		v1.POST("/devices", device.Store)
		v1.PATCH("/devices", device.Update)
		v1.DELETE("/devices/:id", device.Delete)
		v1.GET("/devices/:id", device.Get)
	}
}
