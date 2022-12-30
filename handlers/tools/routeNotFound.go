package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteNotFound(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{
		"message": "Route Does Not Exists",
	})
}
