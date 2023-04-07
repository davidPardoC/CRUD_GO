package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(ctx *gin.Context) {
	users := UserRepositoryImplementation.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}

func PostUserHandler(ctx *gin.Context) {
	var user UserModel
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user = UserRepositoryImplementation.InsertUser(user)
	ctx.JSON(http.StatusOK, user)
}
