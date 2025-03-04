package handler

import (
	"net/http"
	"rupamic-arch/common/middlewares"
	"rupamic-arch/user/contracts"

	"github.com/gorilla/mux"
)

func UserRoutes(svc contracts.ServiceContracts, rl middlewares.RateLimiting) *mux.Router {
	h := newHandler(svc)
	r := mux.NewRouter()
	r.HandleFunc("/user/add", h.AddUser).Methods(http.MethodPost)
	r.Handle("/user/get", middlewares.AuthMiddleware(middlewares.APPIKeyCheck(rl.RateLimiting(http.HandlerFunc(h.GetUser))))).Methods(http.MethodGet)
	// r.HandleFunc("/user/get", h.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user/signin", h.SignInUser).Methods(http.MethodPost)

	return r
}
