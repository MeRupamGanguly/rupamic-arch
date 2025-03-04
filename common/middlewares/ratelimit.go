package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"rupamic-arch/common"
	"strings"

	"github.com/redis/go-redis/v9"
)

type RateLimiting interface {
	RateLimiting(next http.Handler) http.Handler
}
type rlimit struct {
	rdb *redis.Client
}

func NewRLimit(rdb *redis.Client) *rlimit {
	return &rlimit{
		rdb: rdb,
	}
}

func (rl *rlimit) RateLimiting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqIp := strings.Split(r.RemoteAddr, ":")[0]
		queryKey := common.ApiRateLimitKey + reqIp
		fmt.Println("Query Key", queryKey)
		reqCount, err := rl.rdb.Get(context.Background(), queryKey).Int()

		if err == redis.Nil || reqCount < common.ApiRateLimitMaxRequests {
			err = rl.rdb.Incr(context.Background(), queryKey).Err()

			if err == redis.Nil {
				rl.rdb.Expire(context.Background(), queryKey, common.ApiRateLimitDuration)
			}
			fmt.Println("Rate Limiting Called: Count ", reqCount)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, common.ErrRateLimiting.Error(), http.StatusUnauthorized)
		}
	})
}
