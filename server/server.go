package server

import (
	"log"
	"os"
	"os/signal"
	"time"

	"com.thebeachmaster/golangrest/config"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app   *fiber.App
	cfg   *config.Config
	redis *redis.Client
}

func NewServer(cfg *config.Config, cache *redis.Client) *Server {
	return &Server{app: fiber.New(fiber.Config{
		Prefork:      cfg.Server.Prefork,
		ReadTimeout:  time.Second * time.Duration(cfg.Server.ReadTimeout),
		AppName:      cfg.Server.AppName,
		ServerHeader: cfg.Server.ServerHeader,
	}), cfg: cfg, redis: cache}
}

func (srv *Server) Run() error {
	go func() {
		log.Printf("Server is listening on PORT: %s", srv.cfg.Port)
		addr := ":" + srv.cfg.Port
		if err := srv.app.Listen(addr); err != nil {
			log.Panicf("[CRIT] Unable to start server. Reason: %v", err)
		}
	}()

	quitServer := make(chan struct{})

	err := srv.MapHTTPHandlers(srv.app)
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	close(quitServer)

	<-quitServer

	log.Printf("Server shutdown")
	return srv.app.Shutdown()

}
