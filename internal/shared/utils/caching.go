package utils

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetOrCache[T any, U any](
	ctx context.Context,
	client *redis.Client,
	key string,
	ttl time.Duration,
	fetchFunc func() (*T, *U),
) (*T, *U, error) {
	cache, cmdErr := client.Get(ctx, key).Result()
	if cmdErr == nil {
		var result T
		if unmarshalErr := json.Unmarshal([]byte(cache), &result); unmarshalErr == nil {
			return &result, nil, nil
		}
	}
	data, customErr := fetchFunc()
	if customErr != nil {
		return nil, customErr, nil
	}
	bytes, marshalErr := json.Marshal(data)
	if marshalErr == nil {
		cmdErr := client.Set(ctx, key, bytes, ttl)
		if cmdErr.Err() != nil {
			return nil, nil, cmdErr.Err()
		}
	} else {
		return nil, nil, marshalErr
	}
	return data, nil, nil
}

func DeleteCache(
	ctx context.Context,
	client *redis.Client,
	key string,
) error {
	return client.Del(ctx, key).Err()
}
