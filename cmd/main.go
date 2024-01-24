package main

import (
	"log"

	"github.com/Jhon-2801/max-inventory/core/handlers"
	"github.com/Jhon-2801/max-inventory/core/user"
	"github.com/Jhon-2801/max-inventory/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectionDB()

	if err != nil {
		log.Fatalf(err.Error())
	}

	userRepo := user.NewRepo(db)
	userServ := user.NewService(userRepo)
	userEnd := handlers.MakeEndPoints(userServ)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/register", gin.HandlerFunc(userEnd.RegisterUser))

	router.Run(":8080")
}
