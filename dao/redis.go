package dao

import "time"

const RedisKeyPrefix = "liteyuki.captcha:"

func SetStringString(key string, value string) error {
	return redisDB.Set(ctx, RedisKeyPrefix+key, value, 600*time.Second).Err()
}

func GetStringString(key string) (string, error) {
	return redisDB.Get(ctx, RedisKeyPrefix+key).Result()
}
