package user

import (
	"context"

	"com.thebeachmaster/golangrest/internal/user/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.CreateUser) (*models.UserInfo, error)
	Update(ctx context.Context, id string, user *models.CreateUser) (*models.UserInfo, error)
	Read(ctx context.Context, id string) (*models.UserInfo, error)
	Delete(ctx context.Context, id string) (string, error)
}
