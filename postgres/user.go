package postgres

import "gorm.io/gorm"

type User struct {
	UserName string
	gorm.Model
}
