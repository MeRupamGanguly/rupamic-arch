package usecase

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	BtcPrice   = 35600.0
	EthPrice   = 245.0
	SolPrice   = 12.0
	AvaxPrice  = 7.0
	DotPrice   = 0.34
	PriceMutex = &sync.Mutex{}
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func TickProducer(ch chan bool) {
	for {
		select {
		case <-ch:
			log.Println("Tick Producer closed")
			return
		default:
			PriceMutex.Lock()
			BtcPrice += random.Float64()*10 - 5
			EthPrice += random.Float64()*10 - 5
			SolPrice += random.Float64()*1 - 0.5
			AvaxPrice += random.Float64()*1 - 0.5
			DotPrice += random.Float64()*0.05 - 0.025
			PriceMutex.Unlock()
			time.Sleep(time.Second)
		}
	}
}

// "BTCUSDT", "ETHUSDT", "SOLUSDT"
func GetPrice(symbol string) float64 {
	switch symbol {
	case "BTCUSDT":
		return float64(BtcPrice)
	case "ETHUSDT":
		return float64(EthPrice)
	case "SOLUSDT":
		return float64(SolPrice)
	case "AVAXUSDT":
		return float64(AvaxPrice)
	case "DOTUSDT":
		return float64(DotPrice)
	default:
		return 0.0
	}
}
