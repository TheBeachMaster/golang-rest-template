package user

import (
	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(f fiber.Router, u UserHTTPRoutes /*Middlewares*/) {
	f.Post("/new", u.CreateNewUser())
	f.Put("/:id", u.UpdateUser())
	f.Get("/:id", u.GetUserByID())
	f.Delete(":id", u.DeleteUser())
}
