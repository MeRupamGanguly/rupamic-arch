package handler

import (
	"net/http"
	"rupamic-arch/user/contracts"
	"rupamic-arch/user/domain"
)

type handler struct {
	svc contracts.ServiceContracts
}

func newHandler(svc contracts.ServiceContracts) *handler {
	return &handler{
		svc: svc,
	}
}

// AddUser
func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeAddUser(r)
	if err != nil {
		EncodeAddUser(w, "", err)
		return
	}
	id, err := h.svc.AddUser(user)
	if err != nil {
		EncodeAddUser(w, "", err)
		return
	}
	EncodeAddUser(w, id, nil)
}

// GetUser
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, sort, err := DecodeGetUser(r)
	if err != nil {
		EncodeAddUser(w, "", err)
		return
	}
	user, err := h.svc.GetUser(id, sort)
	if err != nil {
		EncodeGetUser(w, domain.User{}, err)
		return
	}
	EncodeGetUser(w, user, nil)
}

// SignInUser
func (h *handler) SignInUser(w http.ResponseWriter, r *http.Request) {
	userId, password, err := DecodeSignIn(r)
	if err != nil {
		EncodeSignIn(w, "", err)
		return
	}
	id, err := h.svc.Signin(userId, password)
	if err != nil {
		EncodeSignIn(w, "", err)
		return
	}
	EncodeSignIn(w, id, nil)
}
