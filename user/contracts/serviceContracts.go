package contracts

import "rupamic-arch/user/domain"

type ServiceContracts interface {
	AddUser(user domain.User) (userId string, err error)
	GetUser(id string) (user domain.User, err error)
	Signin(userid string, password string) (userId string, err error)
}
