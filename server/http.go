package server

import (
	userroute "com.thebeachmaster/golangrest/internal/user"
	delivery "com.thebeachmaster/golangrest/pkg/delivery"
	"github.com/gofiber/fiber/v2"
)

func (srv *Server) MapHTTPHandlers(app *fiber.App) error {

	controller := delivery.New(srv.redis, srv.cfg)

	apiV1 := app.Group("/api/v1")
	userGroup := apiV1.Group("/user")

	userroute.MapUserRoutes(userGroup, controller.NewDelivery().User)

	return nil
}
