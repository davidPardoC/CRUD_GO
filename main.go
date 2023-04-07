package main

import (
	"davidPardoC/gym-app/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database, err := users.ConnectDatabase()
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	database.AutoMigrate(&users.UserModel{})
	userRepo := users.UserRepo{Db: *database}
	users.SetUserRepositoryImplementation(userRepo)
	users.UserRouter(r)

	r.Run()
}
