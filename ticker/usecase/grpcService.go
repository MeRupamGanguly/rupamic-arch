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
	symbols := make(map[string]bool)
	stopChan := make(chan struct{})
	go func() {
		for {
			for s := range symbols {
				resp := &gogen.TickerResponse{
					Symbol:    s,
					Ltp:       GetPrice(s),
					Timestamp: time.Now().Format(time.RFC3339),
				}
				err = stream.Send(resp)
				if err != nil {
					log.Println("Stream Send : ", err)
					close(stopChan)
					return
				}
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Println("Stream Recv : ", err)
			return err
		}
		if _, ok := symbols[req.Symbol]; !ok {
			symbols[req.Symbol] = true
		}
	}
}
