package redis

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yezige/go-short/config"
)

type RedisObject struct {
	c   *redis.Client
	ctx context.Context
}

func New() *RedisObject {
	cfg, _ := config.GetConfig()
	redisConfig := &redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}

	return &RedisObject{
		c:   redis.NewClient(redisConfig),
		ctx: context.Background(),
	}
}

func (r *RedisObject) Set(key string, value interface{}) error {
	return r.c.Set(r.ctx, key, value, 0).Err()
}

func (r *RedisObject) SetTTL(key string, value interface{}, ttl time.Duration) error {
	return r.c.Set(r.ctx, key, value, ttl).Err()
}

func (r *RedisObject) Get(key string) (string, error) {
	val, err := r.c.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key not found")
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func (r *RedisObject) GetDefault(key string, def interface{}) (value interface{}) {
	value, err := r.c.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return def
	} else if err != nil {
		return def
	}

	return value
}

func (r *RedisObject) GetInt64(key string) int64 {
	val, err := r.c.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		return 0
	}
	value, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return value
}

func (r *RedisObject) GetTTL(key string) (time.Duration, error) {
	return r.c.TTL(r.ctx, key).Result()
}
