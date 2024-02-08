package main

import (
	"log"

	HandleRole "github.com/Jhon-2801/max-inventory/core/handlers/roles"
	HandleUser "github.com/Jhon-2801/max-inventory/core/handlers/user"

	"github.com/Jhon-2801/max-inventory/core/roles"
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
	userEnd := HandleUser.MakeEndPoints(userServ)

	roleRepo := roles.NewRepo(db)
	roleServ := roles.NewService(roleRepo)
	rolesEnd := HandleRole.MakeEndPoints(roleServ)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/register", gin.HandlerFunc(userEnd.RegisterUser))
	router.POST("/login", gin.HandlerFunc(userEnd.LoginUser))

	router.POST("/saveRole", gin.HandlerFunc(rolesEnd.SaveUserRole))
	router.DELETE("/delete/:id", gin.HandlerFunc(rolesEnd.RemoveUserRole))

	router.Run(":8080")
}
