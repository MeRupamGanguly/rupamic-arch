package usecase

import (
	"context"
	"fmt"
	"log"
	"rupamic-arch/common"
	"rupamic-arch/ticker/domain/gogen"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcServer(t *testing.T) {
	conn, err := grpc.NewClient(common.GrpcServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		t.Errorf("Test failed: got %v, want %v", err, nil)
	}
	defer conn.Close()
	client := gogen.NewTickerStreamServiceClient(conn)
	stream, err := client.TickerStream(context.Background())
	if err != nil {
		log.Println(err)
		t.Errorf("Test failed: got %v, want %v", err, nil)
	}
	go func() {
		for _, v := range []string{"BTCUSDT", "ETHUSDT", "SOLUSDT"} {
			err = stream.Send(&gogen.TickerRequest{Symbol: v})
			if err != nil {
				log.Println("Test Error Send : ", err)
				t.Errorf("Test failed: got %v, want %v", err, nil)
			}
			time.Sleep(time.Second * 5)
		}
		stream.CloseSend()
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()
	for {
		if ctx.Err() != nil {
			log.Println("Timeout : ", err)
			break
		}
		resp, err := stream.Recv()
		if err != nil {
			log.Println("Test Error Recv : ", err)
			t.Errorf("Test failed: got %v, want %v", err, nil)
			break
		}
		fmt.Println("Streamming : LTP : ", resp.Ltp, " Symbol : ", resp.Symbol)
		time.Sleep(time.Second)
	}
}
