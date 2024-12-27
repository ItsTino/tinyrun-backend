// models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	APIKey   string `gorm:"unique" json:"-"`
}
