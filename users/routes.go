package users

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine) {

	usersRouter := r.Group("/users")
	{
		usersRouter.GET("/", GetUsersHandler)
		usersRouter.POST("/", PostUserHandler)
	}

}
