package shipperorderhandler

import (
	"github.com/waragrammers/Goblue/models"
)

//ShipperOrderRepostory method operation for assigning order to shipper
type ShipperOrderRepostory interface {
	Create() (*models.ShipperOder, error)
	GetOrders() ([]*models.ShipperOder, error)
	GetByID(id uint) (*models.ShipperOder, error)
	DeleteOrder(id uint) (*models.ShipperOder, error)
	UpdateOrder(id uint) (string, error)
}
