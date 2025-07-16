package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	limiterGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memoryStore "github.com/ulule/limiter/v3/drivers/store/memory"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

func NewRateLimiter(rate string) gin.HandlerFunc {
	rateParsed, err := limiter.NewRateFromFormatted(rate)
	if err != nil {
		panic(err)
	}

	store := memoryStore.NewStore()
	return limiterGin.NewMiddleware(limiter.New(store, rateParsed))
}

// Rate limiting using redis
func NewRateLimiterWithRedis(rateStr string) gin.HandlerFunc {
	rate, _ := limiter.NewRateFromFormatted(rateStr)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	store, _ := redisStore.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter",
		MaxRetry: 3,
	})
	return limiterGin.NewMiddleware(limiter.New(store, rate))
}
