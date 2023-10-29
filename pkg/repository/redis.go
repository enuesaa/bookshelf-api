package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"os"
)

type RedisRepositoryInterface interface {
	Keys(pattern string) []string
	// Get(key string) string
	// Set(key string, value string)
	Delete(key string)
	JsonMget(ids []string) [][]byte
	JsonGet(key string) []byte
	JsonSet(key string, value interface{})
	JsonDel(key string)
}

type RedisRepository struct{}

func (repo *RedisRepository) client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	return client
}

func (repo *RedisRepository) jsonHandler() *rejson.Handler {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(repo.client())
	return rh
}

func (repo *RedisRepository) Keys(pattern string) []string {
	vals, _ := repo.client().Keys(context.Background(), pattern).Result()
	return vals
}

func (repo *RedisRepository) Get(key string) string {
	val, err := repo.client().Get(context.Background(), key).Result()
	if err != nil {
		val = ""
	}
	return val
}

func (repo *RedisRepository) Set(key string, value string) {
	err := repo.client().Set(context.Background(), key, value, 0).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func (repo *RedisRepository) Delete(key string) {
	err := repo.client().Del(context.Background(), key).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func (repo *RedisRepository) JsonMget(ids []string) [][]byte {
	data, _ := repo.jsonHandler().JSONMGet(".", ids...)
	if list, ok := data.([]interface{}); ok {
		listbytes := make([][]byte, 0)
		for _, v := range list {
			listbytes = append(listbytes, v.([]byte))
		}
		return listbytes
	}
	return make([][]byte, 0)
}

func (repo *RedisRepository) JsonGet(key string) []byte {
	data, _ := repo.jsonHandler().JSONGet(key, ".")
	return data.([]byte)
}

func (repo *RedisRepository) JsonSet(key string, value interface{}) {
	_, err := repo.jsonHandler().JSONSet(key, ".", value)
	if err != nil {
		fmt.Printf("%+v", err)
	}
}

func (repo *RedisRepository) JsonDel(key string) {
	_, err := repo.jsonHandler().JSONDel(key, "$")
	if err != nil {
		fmt.Printf("%+v", err)
	}
}