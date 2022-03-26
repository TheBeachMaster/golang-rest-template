package routes

import (
	"com.thebeachmaster/golangrest/config"
	"com.thebeachmaster/golangrest/internal/user"
	"com.thebeachmaster/golangrest/internal/user/models"
	"github.com/gofiber/fiber/v2"
)

type userHTTPHandler struct {
	cfg     *config.Config
	usecase user.UserUsecase
}

func NewUserHTTPHandler(c *config.Config, u user.UserUsecase) user.UserHTTPRoutes {
	return &userHTTPHandler{cfg: c, usecase: u}
}

func (u *userHTTPHandler) CreateNewUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &models.CreateUser{}
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		addUser, err := u.usecase.AddNewUser(c.Context(), user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"error":   false,
			"message": nil,
			"user":    addUser,
		})
	}
}

func (u *userHTTPHandler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("id")
		isEmpty := len(userId) == 0

		if isEmpty {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "The param 'id' can not be empty",
			})
		}
		user := &models.CreateUser{}
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		updateUser, err := u.usecase.UpdateUser(c.Context(), userId, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"error":   false,
			"message": nil,
			"user":    updateUser,
		})
	}
}

func (u *userHTTPHandler) GetUserByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("id")
		isEmpty := len(userId) == 0

		if isEmpty {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "The param 'id' can not be empty",
			})
		}

		userInfo, err := u.usecase.GetUserByID(c.Context(), userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"error":   false,
			"message": nil,
			"user":    userInfo,
		})
	}
}

func (u *userHTTPHandler) DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("id")
		isEmpty := len(userId) == 0

		if isEmpty {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "The param 'id' can not be empty",
			})
		}

		deleteUserResult, err := u.usecase.DeleteUser(c.Context(), userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"error":   false,
			"message": nil,
			"status":  deleteUserResult,
		})
	}
}
