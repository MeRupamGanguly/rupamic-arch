package usecase

import (
	"log"
	"rupamic-arch/ticker/domain/gogen"
	"time"
)

type service struct {
	gogen.UnimplementedTickerStreamServiceServer
}

func NewGrpcService() *service {
	return &service{}
}

func (svc *service) TickerStream(stream gogen.TickerStreamService_TickerStreamServer) (err error) {
	req, err := stream.Recv()
	if err != nil {
		log.Println(err)
	}
	resp := gogen.TickerResponse{
		Symbol:    req.Symbol,
		Ltp:       34.89,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	err = stream.Send(&resp)
	if err != nil {
		log.Println(err)
	}
	return
}
