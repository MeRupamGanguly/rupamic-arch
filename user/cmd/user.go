package main

import (
	"log"
	"net/http"
	"rupamic-arch/common"
	"rupamic-arch/common/middlewares"
	"rupamic-arch/user/handler"
	"rupamic-arch/user/repositories"
	"rupamic-arch/user/usecase"

	"github.com/redis/go-redis/v9"
)

func main() {
	logFile := common.SetLogOut()
	defer logFile.Close()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})
	rl := middlewares.NewRLimit(rdb)
	userRepo := repositories.NewUserRepo()
	userSvc := usecase.NewUserService(userRepo)
	r := handler.UserRoutes(userSvc, rl)
	server := http.Server{
		Addr:    common.UserServerPort,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
