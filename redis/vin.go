// Represents DAO layer that interacts with the Redis storage
package redis

import (
	"decoder/models"
	"github.com/go-redis/redis"
)

// VinDAO persists Vin data in database
type VinDAO struct{
	DB *redis.Client
}

// NewVinDAO creates a new VinDAO for Redis implementation
func NewVinDAO(db *redis.Client) *VinDAO {
	return &VinDAO{DB:db}
}

// Read VIN data by pattern from Redis database
func (dao *VinDAO) Get(pattern string) (*models.Vin, error) {
	var vin models.Vin
	return &vin, nil
}