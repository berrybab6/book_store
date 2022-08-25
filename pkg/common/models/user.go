package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User MovieGo
type User struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	ResetToken string `json:"resettoken" gorm:"default:''"`
}

// HashPassword MovieGo
func (user *User) HashPassword(thePassword string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(thePassword), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword MovieGo
func (user *User) CheckPassword(thePassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(thePassword))
	if err != nil {
		return err
	}
	return nil
}
