package router

import (
	"github.com/gin-gonic/gin"
	"juniorconseiltaker-siaje-api/handlers"
	"juniorconseiltaker-siaje-api/handlers/etudiants"
	"juniorconseiltaker-siaje-api/handlers/tools"
)

func Configure(router *gin.Engine) {

	apiV1 := router.Group("/siaje-api/v1")
	{

		apiV1.GET("/health-check", tools.Status)

		apiV1.POST("/authenticate", handlers.Authenticate)

		apiV1.GET("/student/get", etudiants.Get)
		apiV1.POST("/student/create", etudiants.Create)
		apiV1.POST("/student/update", etudiants.Update)
		apiV1.POST("/student/delete", etudiants.Delete)

	}

	router.NoRoute(tools.RouteNotFound)
	router.NoMethod(tools.MethodNotFound)
}
