package usecase

import (
	"context"

	"com.thebeachmaster/golangrest/internal/user"
	"com.thebeachmaster/golangrest/internal/user/models"
)

type userUsecase struct {
	repo user.UserRepository
}

func NewUserUseCase(r user.UserRepository) user.UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) AddNewUser(context context.Context, userInfo *models.CreateUser) (*models.UserInfo, error) {
	return u.repo.Create(context, userInfo)
}

func (u *userUsecase) UpdateUser(ctx context.Context, id string, user *models.CreateUser) (*models.UserInfo, error) {
	return u.repo.Update(ctx, id, user)
}

func (u *userUsecase) GetUserByID(context context.Context, id string) (*models.UserInfo, error) {
	return u.repo.Read(context, id)
}

func (u *userUsecase) DeleteUser(context context.Context, id string) (string, error) {
	return u.repo.Delete(context, id)
}
