package utils

import (
	"fmt"
	"time"

	"github.com/Fuzz-Head/pkg/redisstore"
)

func StoreRefreshToken(userID uint, refreshToken string, expiry time.Duration) error {
	key := fmt.Sprintf("refresh_token:%d", userID)
	return redisstore.Rdb.Set(redisstore.Ctx, key, refreshToken, expiry).Err()
}

func GetRefreshToken(userID uint) (string, error) {
	key := fmt.Sprintf("refresh_token:%d", userID)
	return redisstore.Rdb.Get(redisstore.Ctx, key).Result()
}

func DeleteRefreshToken(userID uint) error {
	key := fmt.Sprintf("refresh_token:%d", userID)
	return redisstore.Rdb.Del(redisstore.Ctx, key).Err()
}
