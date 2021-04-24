package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go_sql_study2/controllers"
)

var userController = controllers.NewUserController()

func Route(router *gin.Engine) {
	v1 := *router.Group("/v1")
	{
		v1.GET("/users", userController.GetUsers)
		v1.POST("/join", userController.Join)
		v1.POST("/update_dt", userController.UpdateLastLoginTime)
	}
}
