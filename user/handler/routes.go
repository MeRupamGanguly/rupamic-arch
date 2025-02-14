package handler

import (
	"net/http"
	"rupamic-arch/user/contracts"

	"github.com/gorilla/mux"
)

func UserRoutes(svc contracts.ServiceContracts) *mux.Router {
	h := newHandler(svc)
	r := mux.NewRouter()
	r.HandleFunc("/user/add", h.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/user/get", h.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user/signin", h.SignInUser).Methods(http.MethodPost)

	return r
}
