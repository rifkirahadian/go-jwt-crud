package models

import(
	"time"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name			string		`json:"name" validate:"required"`
	Description		string		`json:"description" validate:"required"`
	Deadline		time.Time	`json:"deadline" validate:"required"`
}