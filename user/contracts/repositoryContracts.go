package contracts

import "rupamic-arch/user/domain"

type RepositoryContracts interface {
	AddUser(user domain.User) (userId string, err error)
	GetUser(id string, sort string) (user domain.User, err error)
	Signin(userId string) (user domain.User, err error)
}
