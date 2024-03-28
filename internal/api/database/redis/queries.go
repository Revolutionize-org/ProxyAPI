package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/model"
)

func StoreFailedAttempt(ip string) error {
	ctx := context.Background()

	pipe := Instance.Pipeline()

	incrCmd := pipe.Incr(ctx, "limiter:"+ip)

	expireCmd := pipe.Expire(ctx, "limiter:"+ip, 5*time.Minute)

	_, err := pipe.Exec(ctx)

	if err != nil {
		return err
	}

	if incrCmd.Err() != nil {
		return incrCmd.Err()
	}

	if expireCmd.Err() != nil {
		return expireCmd.Err()
	}

	return nil
}

func RetrieveFailedAttempt(ip string) (int64, error) {
	cmd := Instance.Get(context.Background(), "limiter:"+ip)
	attemps, err := cmd.Int64()

	if err != nil {
		return 0, err
	}
	return attemps, nil
}

func StoreApiKey(key *model.ApiKey) error {
	ctx := context.Background()

	expiration := 3 * time.Hour

	err := Instance.HSet(ctx, "api_key:"+key.Key, map[string]interface{}{
		"id":         key.Id,
		"ip_address": key.IpAddress,
	}).Err()

	if err != nil {
		return err
	}

	err = Instance.Expire(ctx, "api_key:"+key.Key, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func RetrieveApiKey(key string) (*model.ApiKey, error) {
	cmd := Instance.HGetAll(context.Background(), "api_key:"+key)

	result := cmd.Val()

	if cmd.Err() == redis.Nil || len(result) == 0 {
		return nil, nil
	}

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	ipAddress := result["ip_address"]
	id := result["id"]

	return &model.ApiKey{
		Id:        id,
		Key:       key,
		IpAddress: ipAddress,
	}, nil
}

func StoreProxy(apiKey *model.ApiKey, proxy string) error {
	ctx := context.Background()

	key := "api_key_proxies:" + apiKey.Id

	if err := Instance.SAdd(ctx, key, proxy).Err(); err != nil {
		return err
	}

	expiration := 3 * time.Hour
	if err := Instance.Expire(ctx, key, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func RetrieveRandomProxy(apiKey *model.ApiKey) (string, error) {
	key := "api_key_proxies:" + apiKey.Id

	cmd := Instance.SRandMember(context.Background(), key)

	if cmd.Err() == redis.Nil || len(cmd.Val()) == 0 {
		return "", nil
	}

	if err := cmd.Err(); err != nil {
		return "", err
	}

	return cmd.Val(), nil
}
