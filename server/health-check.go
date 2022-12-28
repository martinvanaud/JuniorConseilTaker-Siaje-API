package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Status(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{
		"message": "Server is running ...",
	})
}
