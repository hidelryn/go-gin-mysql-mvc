package repositories

import "github.com/go_sql_study2/models"

type UserRepo interface {
	FindByName(user *models.User) (*models.User, error)
	FindByAllUser() ([]*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}
