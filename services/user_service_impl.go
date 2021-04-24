package services

import (
	"errors"

	"github.com/go_sql_study2/models"
	"github.com/go_sql_study2/repositories"
)

type userServiceImpl struct {
	repo  repositories.UserRepo
	users []*models.User
	user  *models.User
	err   error
}

func NewUserService() UserService {
	return &userServiceImpl{
		repo: repositories.UserMySQLRepo(),
	}
}

func (u *userServiceImpl) Join(user *models.User) (*models.User, error) {
	u.err = u.validateDuplicateUser(user)
	if u.err != nil {
		return nil, u.err
	}
	u.user, u.err = u.repo.CreateUser(user)
	if u.err != nil {
		return nil, u.err
	}
	return u.user, nil
}

func (u *userServiceImpl) validateDuplicateUser(user *models.User) error {
	u.user, u.err = u.repo.FindByName(user)
	if u.err != nil {
		return u.err
	}
	if u.user.Id > 0 {
		return errors.New("이미 존재하는 닉네임 입니다")
	}
	return nil
}

func (u *userServiceImpl) GetAllUsers() ([]*models.User, error) {
	u.users, u.err = u.repo.FindByAllUser()
	if u.err != nil {
		return nil, u.err
	}
	return u.users, nil
}

func (u *userServiceImpl) UpdateLastLoginTime(user *models.User) (*models.User, error) {
	u.user, u.err = u.repo.FindByName(user)
	if u.user.Id == 0 {
		return nil, errors.New("존재하지 않은 회원")
	}
	u.user, u.err = u.repo.UpdateUser(u.user)
	if u.err != nil {
		return nil, u.err
	}
	return u.user, nil
}
