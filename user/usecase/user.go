package usecase

import (
	"rupamic-arch/user/contracts"
	"rupamic-arch/user/domain"
)

type service struct {
	db contracts.RepositoryContracts
}

func NewUserService(repo contracts.RepositoryContracts) *service {
	return &service{
		db: repo,
	}
}

func (svc *service) AddUser(user domain.User) (userId string, err error) {
	svc.db.AddUser(user)
	return
}
func (svc *service) GetUser(id string) (user domain.User, err error) {
	svc.db.GetUser(id)
	return
}
func (svc *service) Signin(userid string, password string) (userId string, err error) {
	svc.db.Signin(userId, password)
	return
}
