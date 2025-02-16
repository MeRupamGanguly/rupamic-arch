package repositories

import (
	"log"
	"rupamic-arch/user/domain"
)

type repo struct{}

func NewUserRepo() *repo {
	return &repo{}
}

func (r *repo) AddUser(user domain.User) (userId string, err error) {
	log.Println(user)
	return
}
func (r *repo) GetUser(id string, sort string) (user domain.User, err error) {
	log.Println(id, sort)
	return
}
func (r *repo) Signin(userId string) (user domain.User, err error) {
	log.Println(userId)
	return
}
