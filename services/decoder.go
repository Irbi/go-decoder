// Implements the main business logic of decoding
package services

import "decoder/models"

// VinDAO specifies the interface of the vin DAO needed by DecoderService.
type VinDAO interface {
	// Returns vin data by specified pattern
	Get(pattern string) (*models.Vin, error)
}

// DecoderService provides services related with VIN decoding
type DecoderService struct {
	dao VinDAO
}

// NewDecoderService creates a new DecoderService with the given VIN DAO
func NewDecoderService(dao VinDAO) *DecoderService {
	return &DecoderService{dao}
}

// Return Vin data by the specified pattern.
// In this function must be implemented the business logic of assembling of the raw data into one ordered heap.
// All DAOs must return requested data in the same format
func (s *DecoderService) Get(pattern string) (*models.Vin, error) {
	return s.dao.Get(pattern)
}