package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"rupamic-arch/common"
	"rupamic-arch/user/domain"
)

// AddUser
type AddUserReq struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
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
		Name:     req.Name,
		Email:    req.Email,
		Roles:    req.Roles,
		Password: req.Password,
	}
	return user, nil
}

func EncodeAddUser(w http.ResponseWriter, id string, err error) {
	userResp := AddUserRes{Id: id}
	cResp := common.CustomRespCreator(userResp, err)
	json.NewEncoder(w).Encode(cResp)
}

// GetUser
type GetUserReq struct {
	Id   string `query:"id"`
	Sort string `query:"sort"`
}
type GetUserRes struct {
	User domain.User `json:"user"`
}

func DecodeGetUser(r *http.Request) (id string, sort string, err error) {
	req := GetUserReq{}
	u, err := url.ParseQuery(r.URL.RawQuery)
	req.Id = u["id"][0]
	req.Sort = u["sort"][0]
	if err != nil {
		return "", "", err
	}
	return req.Id, req.Sort, nil
}

func EncodeGetUser(w http.ResponseWriter, user domain.User, err error) {
	getIdResp := GetUserRes{User: user}
	cResp := common.CustomRespCreator(getIdResp, err)
	json.NewEncoder(w).Encode(cResp)
}

// SignIn
type SignInReq struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}
type SignInRes struct {
	Id string `json:"id"`
}

func DecodeSignIn(r *http.Request) (userId string, password string, err error) {
	var req SignInReq
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return "", "", err
	}
	return req.UserId, req.Password, nil
}

func EncodeSignIn(w http.ResponseWriter, id string, err error) {
	signInResp := SignInRes{Id: id}
	cResp := common.CustomRespCreator(signInResp, err)
	json.NewEncoder(w).Encode(cResp)
}
