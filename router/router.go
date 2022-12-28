package router

import (
	"github.com/gin-gonic/gin"
	"juniorconseiltaker-siaje-api/handlers"

	"juniorconseiltaker-siaje-api/middlewares"
	"juniorconseiltaker-siaje-api/server"
)

func Configure(router *gin.Engine) {

	apiV1 := router.Group("/siaje-api/v1")
	{
		apiV1.GET("/health-check/", server.Status)

		apiV1.POST("/authenticate/", handlers.Authenticate)
	}

	router.NoRoute(middlewares.RouteNotFound)
}
