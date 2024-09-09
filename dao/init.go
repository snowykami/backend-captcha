package dao

import (
	"context"
	"git.liteyuki.icu/backend/golite/database"
	"git.liteyuki.icu/backend/golite/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var redisDB *redis.Client
var mysqlDB *gorm.DB
var ctx = context.Background()

func init() {
	redisDB = database.NewRedisClient()
	mysqlDB = database.NewMysqlDB("captcha")
	logger.Info("Init database")
}
