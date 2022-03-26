package delivery

import (
	"com.thebeachmaster/golangrest/internal/user"
	userRepo "com.thebeachmaster/golangrest/internal/user/repository"
	userUsecase "com.thebeachmaster/golangrest/internal/user/usecase"

	userRoutes "com.thebeachmaster/golangrest/internal/user/routes"
)

func (d *deliveryRegistry) NewUserDelivery() user.UserHTTPRoutes {
	repo := userRepo.NewUserRespository(d.redis)
	usecase := userUsecase.NewUserUseCase(repo)
	return userRoutes.NewUserHTTPHandler(d.cfg, usecase)
}
