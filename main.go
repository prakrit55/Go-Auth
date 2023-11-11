package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	db "github.com/prakrit55/Go-Auth/DB"
	user "github.com/prakrit55/Go-Auth/Internal/Users"
)

var r *gin.Engine

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

	InitRouter(userHandler)
	Start(PORT)
}

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
}

func Start(addr string) error {
	return r.Run(addr)
}
