package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Gorm struct {
	Db *gorm.DB
}

func NewGorm() Gorm {
	dsn := "host=localhost user=postgres password=postgres dbname=ddd_practice port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("fail, run postgres open\nerr: %e", err)
	}

	return Gorm{db}
}
