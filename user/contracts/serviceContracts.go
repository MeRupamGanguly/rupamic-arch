package contracts

import "rupamic-arch/user/domain"

type ServiceContracts interface {
	AddUser(user domain.User) (userId string, err error)
	GetUser(id string, sort string) (user domain.User, err error)
	Signin(userId string, password string) (user domain.User, err error)
}
