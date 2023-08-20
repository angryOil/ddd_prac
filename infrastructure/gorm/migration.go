package gorm

import "ddd_prac/infrastructure/user/model"

func (g Gorm) MigrateUserTable() error {
	return g.Db.AutoMigrate(&model.Users{})
}
