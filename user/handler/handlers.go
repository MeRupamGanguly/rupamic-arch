package handler

import (
	"net/http"
	"rupamic-arch/user/contracts"
)

type handler struct {
	svc contracts.ServiceContracts
}

func newHandler(svc contracts.ServiceContracts) *handler {
	return &handler{
		svc: svc,
	}
}
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
