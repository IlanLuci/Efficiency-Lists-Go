package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID uint   `gorm:"autoIncrement"`
	Data   string `gorm:"type:text"`
}
