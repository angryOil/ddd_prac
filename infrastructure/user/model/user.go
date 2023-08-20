package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email string
	Pw    string
}
