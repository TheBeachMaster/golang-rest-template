package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"com.thebeachmaster/golangrest/internal/user"
	"com.thebeachmaster/golangrest/internal/user/models"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type RedisUserData struct {
	FirstName    string
	LastName     string
	Username     string
	EmailAddress string
	Password     string
	UserID       string
}

const (
	REDIS_USER_KEY_PREFIX = "user"
	REDIS_USER_FIELD      = "user"
)

type userRepository struct {
	client *redis.Client
}

func NewUserRespository(client *redis.Client) user.UserRepository {
	return &userRepository{client: client}
}

func (u *userRepository) Create(ctx context.Context, user *models.CreateUser) (*models.UserInfo, error) {
	userId := uuid.New().String()
	dbData := &RedisUserData{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		EmailAddress: user.EmailAddress,
		Password:     user.Password,
		UserID:       userId,
	}
	redisKey := generateKey(REDIS_USER_KEY_PREFIX, userId)
	err := u.client.HSetNX(ctx, redisKey, REDIS_USER_FIELD, dbData).Err()
	if err != nil {
		return nil, err
	}
	userInfo := &models.UserInfo{
		FirstName:    dbData.FirstName,
		LastName:     dbData.LastName,
		Username:     dbData.Username,
		EmailAddress: dbData.EmailAddress,
		UserID:       dbData.UserID,
	}
	return userInfo, nil
}

func (u *userRepository) Update(ctx context.Context, id string, user *models.CreateUser) (*models.UserInfo, error) {
	lookupKey := generateKey(REDIS_USER_KEY_PREFIX, id)
	_, err := u.client.HGet(ctx, lookupKey, REDIS_USER_FIELD).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("unknown user ID")
	}
	if err != nil {
		return nil, err
	}
	dbData := &RedisUserData{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		EmailAddress: user.EmailAddress,
		Password:     user.Password,
		UserID:       id,
	}
	err = u.client.HSetNX(ctx, lookupKey, REDIS_USER_FIELD, dbData).Err()
	if err != nil {
		return nil, err
	}
	userInfo := &models.UserInfo{
		FirstName:    dbData.FirstName,
		LastName:     dbData.LastName,
		Username:     dbData.Username,
		EmailAddress: dbData.EmailAddress,
		UserID:       dbData.UserID,
	}
	return userInfo, nil
}

func (u *userRepository) Read(ctx context.Context, id string) (*models.UserInfo, error) {
	lookupKey := generateKey(REDIS_USER_KEY_PREFIX, id)
	lookupResult, err := u.client.HGet(ctx, lookupKey, REDIS_USER_FIELD).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("unknown user ID")
	}
	if err != nil {
		return nil, err
	}
	userData := &RedisUserData{}
	if err := userData.UnMarshalBinary([]byte(lookupResult)); err != nil {
		return nil, fmt.Errorf("unable to unmarshall data from cache")
	}
	userInfo := &models.UserInfo{
		FirstName:    userData.FirstName,
		LastName:     userData.LastName,
		Username:     userData.Username,
		EmailAddress: userData.EmailAddress,
		UserID:       userData.UserID,
	}
	return userInfo, nil
}

func (u *userRepository) Delete(ctx context.Context, id string) (string, error) {
	redisKey := generateKey(REDIS_USER_KEY_PREFIX, id)
	_, err := u.client.HDel(ctx, redisKey, REDIS_USER_FIELD).Result()
	if err != nil {
		return "", err
	}
	return "deleted", nil
}

func generateKey(prefix string, key string) string {
	keyBuilder := strings.Builder{}
	keyBuilder.WriteString(fmt.Sprintf("%s:%s", prefix, key))
	return keyBuilder.String()
}

func (m *RedisUserData) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *RedisUserData) UnMarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}
