package main

import (
	"github.com/joho/godotenv"

	"juniorconseiltaker-siaje-api/router"
	"juniorconseiltaker-siaje-api/server"

	"log"
	"os"
)

func main() {

	errorEnv := godotenv.Load("docker/.env")
	if errorEnv != nil {
		log.Fatalf("godotenv: could not properly setup instance %v", errorEnv)
		return
	}

	//hrDatabase, errorDatabase := database.Configure()
	//if errorDatabase != nil {
	//	log.Fatalf("database: could not properly setup instance %v", errorDatabase)
	//}
	//
	//defer func(hrDatabase *sql.DB) {
	//	err := hrDatabase.Close()
	//	if err != nil {
	//
	//	}
	//}(hrDatabase)

	hrServer := server.InitServer()

	router.Configure(hrServer.Def)

	errorServer := hrServer.Def.Run(":" + os.Getenv("API_PORT"))
	if errorServer != nil {
		log.Fatalf("hrServer.run: could not proerly setup instance %v", errorServer)
	}

}
