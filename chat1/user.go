package main

import (
	"fmt"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type FUsers struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	PasswordH string `json:"password"`
}

type uService struct {
}

var DefaultService uService

func (s *uService) FindUser(newUser *Users) error {
	var x string
	var t error
	x, t = s.CheckUser(newUser)
	if (t != nil){
		return t
	}
	fmt.Println(x)
	err := bcrypt.CompareHashAndPassword([]byte(x), []byte(newUser.Password))
	return err
}

func (s *uService) CreateUser(newUser *Users) error {
	_, t := s.CheckUser(newUser)
	if (t==nil){
		return errors.New("Email Already Exists")
	}
	ph, err := PasswordHash(newUser.Password)
	if err != nil {
		return errors.New("PasswordHash weird")
	}
	newFUser := FUsers{
		Email:     newUser.Email,
		Username:  newUser.Username,
		PasswordH: ph,
	}
	return s.InsertUser(&newFUser)
}
func PasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}
