package user

import (
	"context"
	"encoding/json"

	"github.com/Abdelrhmanfdl/user-service/internal/models"
	"github.com/go-redis/redis/v8"
)

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository(redisURL string) *RedisUserRepository {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})
	return &RedisUserRepository{client: redisClient}
}

func (repo *RedisUserRepository) GetUser(id string) (user models.User, err error) {
	value, err := repo.client.Get(context.Background(), id).Result()
	if err != nil {
		return models.User{}, err
	}

	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repo *RedisUserRepository) SaveUser(user models.User) (err error) {
	return repo.client.Set(context.Background(), repo.FormatKey(user.ID), user, 0).Err()
}

func (repo *RedisUserRepository) FormatKey(id string) string {
	return "user:" + id
}
