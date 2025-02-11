package main

import (
	"log"
	"net/http"
	"rupamic-arch/common"
	"rupamic-arch/user/handler"
	"rupamic-arch/user/repositories"
	"rupamic-arch/user/usecase"
)

func main() {
	logFile := common.SetLogOut()
	defer logFile.Close()

	userRepo := repositories.NewUserRepo()
	userSvc := usecase.NewUserService(userRepo)
	r := handler.UserRoutes(userSvc)
	server := http.Server{
		Addr:    common.UserServerPort,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
