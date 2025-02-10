package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"rupamic-arch/common"
	"rupamic-arch/ticker/domain/gogen"
	"rupamic-arch/ticker/usecase"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	logFile := common.SetLogOut()
	logFile.Close()
	termSig := make(chan os.Signal, 1)
	signal.Notify(termSig, syscall.SIGINT, syscall.SIGTERM)
	closeCh := make(chan bool)
	go usecase.TickProducer(closeCh)

	grpcServer := grpc.NewServer()
	svc := usecase.NewGrpcService()
	gogen.RegisterTickerStreamServiceServer(grpcServer, svc)
	lis, err := net.Listen("tcp", common.GrpcServerPort)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	<-termSig
	closeCh <- true
	grpcServer.GracefulStop()
	log.Println("stoped")
}
