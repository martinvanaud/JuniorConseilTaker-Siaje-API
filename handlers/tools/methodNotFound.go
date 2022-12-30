package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MethodNotFound(context *gin.Context) {
	context.JSON(http.StatusMethodNotAllowed, gin.H{
		"message": "Method Does Not Exists",
	})
}
