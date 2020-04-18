package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name		string	`json:"name" validate:"required"`
	Email		string	`json:"email" validate:"required,email" gorm:"type:varchar(100);unique_index"	`
	Password	string	`json:"password" validate:"required"`
}