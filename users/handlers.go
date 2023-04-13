package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
}

func GetUsersHandler(ctx *gin.Context) {
	users := UserRepositoryImplementation.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}

func PostUserHandler(ctx *gin.Context) {

	var body []byte

	ctx.Request.Body.Read(body)

	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword
	user, err = UserRepositoryImplementation.InsertUser(user)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	userResponse := UserResponse{Email: user.Email, Id: user.Id}

	ctx.JSON(http.StatusOK, userResponse)
}
