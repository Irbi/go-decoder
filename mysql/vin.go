// Represents DAO layer that interacts with the MySQL storage
package mysql

import (
	"decoder/models"
	"database/sql"
)

// VinDAO persists Vin data in database
type VinDAO struct{
	DB *sql.DB
}

// NewVinDAO creates a new VinDAO
func NewVinDAO(db *sql.DB) *VinDAO {
	return &VinDAO{DB: db}
}

// Read VIN data by pattern from MySQL database. Use VinDao:DB as connection handler
func (dao *VinDAO) Get(pattern string) (*models.Vin, error) {
	var vin models.Vin
	return &vin, nil
}