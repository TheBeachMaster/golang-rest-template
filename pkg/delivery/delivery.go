package delivery

import (
	"com.thebeachmaster/golangrest/config"
	"com.thebeachmaster/golangrest/internal/delivery"
	"github.com/go-redis/redis/v8"
)

type deliveryRegistry struct {
	redis *redis.Client
	cfg   *config.Config
}

type DeliveryRegistry interface {
	NewDelivery() delivery.Delivery
}

func New(redis *redis.Client, cfg *config.Config) DeliveryRegistry {
	return &deliveryRegistry{
		redis: redis,
		cfg:   cfg,
	}
}

func (d *deliveryRegistry) NewDelivery() delivery.Delivery {
	return delivery.Delivery{
		User: d.NewUserDelivery(),
	}
}
