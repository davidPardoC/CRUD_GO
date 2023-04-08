package main

import (
	"davidPardoC/gym-app/database"
	"davidPardoC/gym-app/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	users.UserRouter(r)
}

func SetupRepository() error {
	database, err := database.GetDatabase()
	database.AutoMigrate(&users.User{})
	if err != nil {
		return err
	}
	userRepo := users.UserRepo{Db: *database}
	users.SetUserRepositoryImplementation(userRepo)
	return nil
}
