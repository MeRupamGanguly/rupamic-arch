package handler

import (
	"encoding/json"
	"net/http"
	"rupamic-arch/common"
	"rupamic-arch/user/domain"
)

type AddUserReq struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}
type AddUserRes struct {
	Id string `json:"id"`
}

func DecodeAddUser(r *http.Request) (domain.User, error) {
	var req AddUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		Name:  req.Name,
		Email: req.Email,
		Roles: req.Roles,
	}
	return user, nil
}

func EncodeAddUser(w http.ResponseWriter, id string, err error) {
	userResp := AddUserRes{Id: id}
	cResp := common.CustomRespCreator(userResp, err)
	json.NewEncoder(w).Encode(cResp)
}
