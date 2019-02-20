package oderrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/controllers/orderhandler"
	"github.com/waragrammers/Goblue/models"
)

//OrderRepo for brouth the database in you arsenal
type OrderRepo struct {
	DbConnection *gorm.DB
}

// NewOrderRepo for initialize new orderrepo object
func NewOrderRepo(DbCon *gorm.DB) orderhandler.OrderRepostory {
	return &OrderRepo{DbCon}
}

//CreateOrder for saving the order in database
func (db *OrderRepo) CreateOrder() (*models.Order, error) {
	//database access
	return nil, nil
}

//GetOrderByID for get single orderInfo in database
func (db *OrderRepo) GetOrderByID(id uint) (*models.Order, error) {

	//database access
	return nil, nil
}

//GetOrders get all ordered orders in database
func (db *OrderRepo) GetOrders() ([]*models.Order, error) {
	//database access
	return nil, nil
}

//DeleteOrder for delete single order record in database
func (db *OrderRepo) DeleteOrder(id uint) (*models.Order, error) {
	//database access
	return nil, nil
}

//UpdateOrder FOR UPDATING order information in database
func (db *OrderRepo) UpdateOrder(id uint) (string, error) {
	message := "Update successfuly"
	return message, nil
}
