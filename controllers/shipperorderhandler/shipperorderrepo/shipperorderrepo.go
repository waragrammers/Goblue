package shipperorderrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/controllers/shipperorderhandler"
	"github.com/waragrammers/Goblue/models"
)

//ShipperOrderrepo  for brouth the database in you arsenal
type ShipperOrderrepo struct {
	DbConnection *gorm.DB
}

// NewShipperOrderrepo for initialize new orderrepo object
func NewShipperOrderrepo(DbCon *gorm.DB) shipperorderhandler.ShipperOrderRepostory {
	return &ShipperOrderrepo{DbCon}
}

//Create for assign order in to shipper
func (db *ShipperOrderrepo) Create() (*models.ShipperOder, error) {
	//database access
	return nil, nil
}

//GetByID for get single orderInfo in database
func (db *ShipperOrderrepo) GetByID(id uint) (*models.ShipperOder, error) {

	//database access
	return nil, nil
}

//GetOrders get all ordered orders in database
func (db *ShipperOrderrepo) GetOrders() ([]*models.ShipperOder, error) {
	//database access
	return nil, nil
}

//DeleteOrder for delete single order record in database
func (db *ShipperOrderrepo) DeleteOrder(id uint) (*models.ShipperOder, error) {
	//database access
	return nil, nil
}

//UpdateOrder FOR UPDATING order information in database
func (db *ShipperOrderrepo) UpdateOrder(id uint) (string, error) {
	message := "Update successfuly"
	return message, nil
}
