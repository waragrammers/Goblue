package orderhandler

import (
	"github.com/waragrammers/Goblue/models"
)

//OrderRepostory method operation for order
type OrderRepostory interface {
	CreateOrder() (*models.Order, error)
	GetOrders() ([]*models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	DeleteOrder(id uint) (*models.Order, error)
	UpdateOrder(id uint) (string, error)
}
