package models

import (
	"choreboard/utils"
	"errors"
	"time"
)

//interface for service object that interacts with model and db
type UserService interface {
	GetUser(userid string) (User, error)
	UpdateUser(user *User) error
	CreateUser(user *User) error
	LoginUser(email string, password string) (*User, error)
}

//model
type User struct {
	Id        string     `json:"id" gorethink:"id,omitempty"`
	FirstName string     `json:"first_name" gorethink:"first_name,omitempty" binding:"required"`
	LastName  string     `json:"last_name" gorethink:"last_name,omitempty" binding:"required"`
	Email     string     `json:"email" gorethink:"email,omitempty" binding:"required"`
	Password  string     `json:"-" gorethink:"password,omitempty" binding:"required"`
	Created   *time.Time `json:"created" gorethink:"created,omitempty"`
}

func (u *User) SetEncryptPassword(password string) error {
	encryptedPassword, err := utils.Encrypt([]byte(password))

	if err != nil {
		return errors.New("Could not save password")
	}

	u.Password = utils.EncodeBase64(encryptedPassword)

	return nil
}

func (u *User) GetDecryptPassword() (string, error) {
	pwd, err := utils.Decrypt(utils.DecodeBase64(u.Password))

	if err != nil {
		return "", errors.New("Could not decrypt password")
	}

	return pwd, nil
}
