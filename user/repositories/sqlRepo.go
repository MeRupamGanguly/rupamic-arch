package repositories

import "rupamic-arch/user/domain"

type repo struct{}

func NewUserRepo() *repo {
	return &repo{}
}

func (r *repo) AddUser(user domain.User) (userId string, err error) {
	return
}
func (r *repo) GetUser(id string) (user domain.User, err error) {
	return
}
func (r *repo) Signin(userid string, password string) (userId string, err error) {
	return
}
