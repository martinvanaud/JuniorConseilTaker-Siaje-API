package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"os"
)

type Server struct {
	Def *gin.Engine
}

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func InitServer() (Server) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, PUT, POST, PATCH, DELETE, OPTION"},
		AllowHeaders:     []string{"Origin, Authorization, Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		AllowWebSockets:  false,
	}))

	server := Server{
		Def: router,
	}

	return server
}
