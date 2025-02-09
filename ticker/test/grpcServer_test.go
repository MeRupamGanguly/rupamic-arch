package test

import (
	"context"
	"fmt"
	"log"
	"rupamic-arch/common"
	"rupamic-arch/ticker/domain/gogen"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcServer(t *testing.T) {
	conn, err := grpc.NewClient(common.GrpcServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	client := gogen.NewTickerStreamServiceClient(conn)
	stream, err := client.TickerStream(context.Background())
	if err != nil {
		log.Println(err)
	}
	stream.Send(&gogen.TickerRequest{Symbol: "BTCUSDT"})
	resp, err := stream.Recv()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp.Ltp, resp.Symbol)
}
