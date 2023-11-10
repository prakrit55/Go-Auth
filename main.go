package main

import (
	"log"
	"os"

	db "github.com/prakrit55/Go-Chat/DB"
	user "github.com/prakrit55/Go-Chat/Internal/Users"
	router "github.com/prakrit55/Go-Chat/Router"
)

func main() {
	// Creates a new database
	dbConn, err := db.NewChatDB()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepo(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	router.InitRouter(userHandler)
	router.Start(PORT)
}
