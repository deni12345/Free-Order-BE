package routes

import (
	"github/lambda-microservice/api"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/getUser", api.GetUser)
		user.GET("/createUser", api.CreateUser)
	}
}
