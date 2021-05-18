package scooter

import (
	"legendaryenchilada/model"
	"legendaryenchilada/service"
)

// declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type scooterService interface {
	GetScootersNearby(lat float64, lng float64, distance int64, limit int64) ([]model.Point, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type Controller struct {
	scooterService
}

// InitController initializes the user controller.
func InitController(scooterService *service.ScooterService) *Controller {
	return &Controller{
		scooterService,
	}
}

func (c *Controller) GetNearbyScooters(lat float64, lng float64, distance int64, limit int64) ([]model.Point, error) {
	scooters, err := c.scooterService.GetScootersNearby(lat, lng, distance, limit)
	return scooters, err
}
