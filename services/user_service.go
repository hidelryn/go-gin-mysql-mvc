package services

import "github.com/go_sql_study2/models"

type UserService interface {
	Join(user *models.User) (*models.User, error)
	validateDuplicateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error)
	UpdateLastLoginTime(user *models.User) (*models.User, error)
}
