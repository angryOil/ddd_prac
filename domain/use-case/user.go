package use_case

import (
	"ddd_prac/port/out/repository/user"
	"fmt"
)

type LoginUser struct {
	ID int64
}

func (l LoginUser) ToString() string {
	return fmt.Sprintf("userId: %v", l.ID)
}

type UserServiceImpl struct {
	repository user.Repository
}

func NewUserService(repository user.Repository) UserService {
	return &UserServiceImpl{repository: repository}
}

func (u UserServiceImpl) Login(email, pw string) (l *LoginUser, err error) {
	id, err := u.repository.SearchUser(email, pw)
	if err != nil {
		return nil, err
	}
	return &LoginUser{ID: id}, nil
}

func (u UserServiceImpl) Register(email, pw string) (err error) {
	return u.repository.CreateUser(email, pw)
}
