package users

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine) {

	usersRouter := r.Group("api/v1/users")
	{
		usersRouter.GET("/", GetUsersHandler)
		usersRouter.POST("/", PostUserHandler)
	}

}
