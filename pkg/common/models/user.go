package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	ResetToken string `json:"resettoken" gorm:"default:''"`
}

func (user *User) HashPassword(thePassword string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(thePassword), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(thePassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(thePassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GenerateResetToken() string {

	return ""
}
