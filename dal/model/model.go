package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        // ID uint CreatAt time.Time UpdateAt time.Time DeleteAt gorm.DeleteAt If it is repeated with the definition will be ignored
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"column:name"`
}
