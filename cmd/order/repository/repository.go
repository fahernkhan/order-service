package repository

import (
	// external package
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Database *gorm.DB
	Redis    *redis.Client
}

// NewOrderRepository new order repository by given db pointer of gorm.DB, and redis pointer of redis.Client.
//
// It returns pointer of OrderRepository when successful.
// Otherwise, nil pointer of OrderRepository will be returned.
func NewOrderRepository(db *gorm.DB, redis *redis.Client) *OrderRepository {
	return &OrderRepository{
		Database: db,
		Redis:    redis,
	}
}
