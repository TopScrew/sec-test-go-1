package db

/**
 * @Author: Screw
 * @Date: 2021/1/8 11:42 下午
 * @Desc:
 */
import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	Client *redis.Client
	Prefix string
}

func NewRedis(redisConfig string) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redisConfig.Address",
		Password: "redisConfig.Password",
		DB:       5,
	})
	pong, err := client.Ping().Result()
	if pong != "PONG" || err != nil {
		return nil, err
	}
	return &Redis{client, "redisConfig.Prefix"}, nil
}

func (r *Redis) Get(key string) (string, error) {
	value, err := r.Client.Get(fmt.Sprintf("%s%s", r.Prefix, key)).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return value, nil
}

func (r *Redis) MGet(keys []string) ([]string, error) {
	for i, key := range keys {
		keys[i] = fmt.Sprintf("%s%s", r.Prefix, key)
	}
	results, err := r.Client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}
	values := []string{}
	for _, result := range results {
		if result != nil {
			values = append(values, result.(string))
		} else {
			values = append(values, "")
		}
	}
	return values, nil
}

func (r *Redis) Set(key string, value string, expire time.Duration) error {
	return r.Client.Set(fmt.Sprintf("%s%s", r.Prefix, key), value, expire).Err()
}

func (r *Redis) SetNX(key string, value string, expire time.Duration) (bool, error) {
	return r.Client.SetNX(fmt.Sprintf("%s%s", r.Prefix, key), value, expire).Result()
}

func (r *Redis) MSet(values map[string]string, expire time.Duration) error {
	pipe := r.Client.Pipeline()
	for key, value := range values {
		pipe.Set(fmt.Sprintf("%s%s", r.Prefix, key), value, expire)
	}
	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Expire(key string, expire time.Duration) error {
	return r.Client.Expire(fmt.Sprintf("%s%s", r.Prefix, key), expire).Err()
}

func (r *Redis) SAdd(key string, value string, expire time.Duration) (int64, error) {
	result, err := r.Client.SAdd(fmt.Sprintf("%s%s", r.Prefix, key), value).Result()
	if err != nil {
		return result, err
	}
	var expireLast time.Duration
	expireLast, err = r.Client.TTL(fmt.Sprintf("%s%s", r.Prefix, key)).Result()
	if expire > expireLast {
		r.Client.Expire(fmt.Sprintf("%s%s", r.Prefix, key), expire)
	}
	return result, nil
}

func (r *Redis) SIsMember(key string, value string) (bool, error) {
	return r.Client.SIsMember(fmt.Sprintf("%s%s", r.Prefix, key), value).Result()
}

func (r *Redis) SMembers(key string) ([]string, error) {
	return r.Client.SMembers(fmt.Sprintf("%s%s", r.Prefix, key)).Result()
}

func (r *Redis) SUnion(keys []string) ([]string, error) {
	for index, key := range keys {
		keys[index] = fmt.Sprintf("%s%s", r.Prefix, key)
	}
	return r.Client.SUnion(keys...).Result()
}

func (r *Redis) SRem(key string, value string) (int64, error) {
	return r.Client.SRem(fmt.Sprintf("%s%s", r.Prefix, key), value).Result()
}

func (r *Redis) TokenGet(key string) (string, error) {
	value, err := r.Client.Get(fmt.Sprintf("%s%s", r.Prefix, key)).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (r *Redis) Del(key string) error {
	return r.Client.Del(fmt.Sprintf("%s%s", r.Prefix, key)).Err()
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
