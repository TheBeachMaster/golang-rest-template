package user

import "github.com/gofiber/fiber/v2"

type UserHTTPRoutes interface {
	CreateNewUser() fiber.Handler
	UpdateUser() fiber.Handler
	GetUserByID() fiber.Handler
	DeleteUser() fiber.Handler
}
