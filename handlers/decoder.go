// API layer of the decoding endpoint with the only one responsibility:
// push request to the corresponding service and return the requested data to the client
package handlers

import (
	"github.com/go-ozzo/ozzo-routing"
	"decoder/models"
)

type (
	// decoderService specifies the interface for the service needed by decoderResource
	// Add here more methods for CRUD
	decoderService interface {
		Get(pattern string) (*models.Vin, error)
	}

	// decoderResource defines the handlers for the CRUD APIs.
	decoderResource struct {
		service decoderService
	}
)

// ServeDecoderResource sets up the routing of decode endpoints and the corresponding handlers
// Add here more methods for CRUD
func ServeDecoderResource(rg *routing.RouteGroup, service decoderService) {
	resource := &decoderResource{service}

	rg.Get("/decode/<id>", resource.get)
}

// Handle '/decode/<id>' requests
func (resource *decoderResource) get(c *routing.Context) error {
	pattern := c.Param("pattern")

	response, err := resource.service.Get(pattern)
	if err != nil {
		return err
	}

	return c.Write(response)
}
