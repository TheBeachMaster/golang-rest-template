package user

import (
	"context"

	"com.thebeachmaster/golangrest/internal/user/models"
)

type UserUsecase interface {
	AddNewUser(context context.Context, userInfo *models.CreateUser) (*models.UserInfo, error)
	UpdateUser(ctx context.Context, id string, user *models.CreateUser) (*models.UserInfo, error)
	GetUserByID(context context.Context, id string) (*models.UserInfo, error)
	DeleteUser(context context.Context, id string) (string, error)
}
