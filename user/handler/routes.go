package handler

import (
	"rupamic-arch/user/contracts"

	"github.com/gorilla/mux"
)

func UserRoutes(svc contracts.ServiceContracts) *mux.Router {
	h := newHandler(svc)
	r := mux.NewRouter()
	r.HandleFunc("/user/add", h.AddUser)

	return r
}
