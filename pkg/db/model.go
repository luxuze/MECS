package db

import "github.com/jinzhu/gorm"

type Code struct {
	gorm.Model
	Code  string `gorm:"unique"`
	Proxy string
	Local string
}
