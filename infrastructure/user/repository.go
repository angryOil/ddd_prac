package user

import (
	"ddd_prac/infrastructure/gorm"
	"ddd_prac/infrastructure/user/model"
	"ddd_prac/port/out/repository/user"
	"errors"
)

type RepositoryImpl struct {
	gorm gorm.Gorm
}

func (r RepositoryImpl) SearchUser(email, pw string) (int64, error) {
	return 999, nil
}

func NewRepositoryImpl(gorm gorm.Gorm) user.Repository {
	return RepositoryImpl{gorm: gorm}
}

func (r RepositoryImpl) CreateUser(email, pw string) error {

	var duplicateUser model.Users
	r.gorm.Db.Where("email = ?", email).
		First(&duplicateUser)

	if duplicateUser.ID != 0 {
		return errors.New("duplicate email for: " + email)
	}
	var u = model.Users{
		Email: email,
		Pw:    pw,
	}
	r.gorm.Db.Create(&u)
	return nil
}
