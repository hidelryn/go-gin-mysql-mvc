package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go_sql_study2/libs"
	"github.com/go_sql_study2/models"
	"github.com/go_sql_study2/services"
)

type UserController struct {
	userService services.UserService
	user        *models.User
	users       []*models.User
	err         error
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (u *UserController) Join(c *gin.Context) {
	var user models.User
	err := libs.RequestParams(c, &user)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	u.user, u.err = u.userService.Join(&user)
	libs.ResponseJSON(c, u.err, u.user)
}

func (u *UserController) GetUsers(c *gin.Context) {
	u.users, u.err = u.userService.GetAllUsers()
	libs.ResponseJSON(c, u.err, u.users)
}

func (u *UserController) UpdateLastLoginTime(c *gin.Context) {
	var user models.User
	err := libs.RequestParams(c, &user)
	if err != nil {
		return
	}
	u.user, u.err = u.userService.UpdateLastLoginTime(&user)
	libs.ResponseJSON(c, u.err, u.user)
}
