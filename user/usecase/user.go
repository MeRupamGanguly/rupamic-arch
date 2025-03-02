package usecase

import (
	"rupamic-arch/common"
	"rupamic-arch/user/contracts"
	"rupamic-arch/user/domain"
	"strings"
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
	var roles []string
	if !strings.Contains(user.Email, "tm.com") {
		for _, role := range user.Roles {
			if role != common.ADMINROLE || role != common.SUPERADMINROLE {
				roles = append(roles, role)
			}
		}
		user.Roles = roles
	}
	svc.db.AddUser(user)
	return
}
func (svc *service) GetUser(id string, sort string) (user domain.User, err error) {
	svc.db.GetUser(id, sort)
	return
}
func (svc *service) Signin(userId string, password string) (user domain.User, err error) {
	user, err = svc.db.Signin(userId)
	if err != nil {
		return
	}
	matched, err := common.Decrypt(user.Password, password)
	if !matched {
		return domain.User{}, common.ErrUserCredWrong
	}
	return
}
