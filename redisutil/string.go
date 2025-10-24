package redisutil

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// SetString 设置字符串值
func SetString(c *redis.Client, ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.Set(ctx, key, value, expiration).Err()
}

// GetString 获取字符串值
func GetString(c *redis.Client, ctx context.Context, key string) (string, error) {
	val, err := c.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return val, err
}

// DelKey 删除键
func DelKey(c *redis.Client, ctx context.Context, keys ...string) error {
	return c.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func Exists(c *redis.Client, ctx context.Context, key string) (bool, error) {
	n, err := c.Exists(ctx, key).Result()
	return n > 0, err
}

// Expire 设置过期时间
func Expire(c *redis.Client, ctx context.Context, key string, expiration time.Duration) error {
	return c.Expire(ctx, key, expiration).Err()
}

// GetUint64 获取无符号64位整数值
func GetUint64(c *redis.Client, ctx context.Context, key string) (uint64, error) {
	val, err := c.Do(ctx, "GET", key).Uint64()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	return val, err
}

// SetStruct 将结构体序列化为JSON存储到Redis
func SetStruct(c *redis.Client, ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, string(data), expiration).Err()
}

// GetStruct 从Redis获取JSON并反序列化到结构体
func GetStruct(c *redis.Client, ctx context.Context, key string, dest interface{}) error {
	data, err := c.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// key 不存在，不报错
		return nil
	}
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal([]byte(data), dest)
}
