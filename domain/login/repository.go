package login

import "fmt"

type Repository interface {
	PerformLogin(*Login) error
}

type loginRepository struct{}

func (loginRepo *loginRepository) PerformLogin(u *Login) error{
	fmt.Println(u.Username)
	return nil
}