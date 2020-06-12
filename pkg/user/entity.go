package user

import "github.com/jinzhu/gorm"

//User model for clean notes api
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
