package models

import "github.com/jinzhu/gorm"

type Auth struct {
	gorm.Model
	Username string `grom:"type:varchar(50);default:''"json:"username"`
	Password string `gorm:"type:varchar(50);default:''"json:"password"`
}

func CheckAuth(usename, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: usename,Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
