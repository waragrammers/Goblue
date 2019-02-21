package sellerrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/controllers/sellerhandler"
	"github.com/waragrammers/Goblue/models"
)

//SellerRepo for brougth the database in you arsenal
type SellerRepo struct {
	DbConnection *gorm.DB
}

// NewSellerRepo for initialize new orderrepo object
func NewSellerRepo(DbCon *gorm.DB) sellerhandler.SellerRepository {
	return &SellerRepo{DbCon}
}

//Create for assign order in to shipper
func (db *SellerRepo) Create() (*models.Seller, error) {
	//database access
	return nil, nil
}

//GetByID for get single orderInfo in database
func (db *SellerRepo) GetByID(id uint) (*models.Seller, error) {

	//database access
	return nil, nil
}

//FechAll get all ordered orders in database
func (db *SellerRepo) FechAll() ([]*models.Seller, error) {
	//database access
	return nil, nil
}

//Delete for delete single order record in database
func (db *SellerRepo) Delete(id uint) (*models.Seller, error) {
	//database access
	return nil, nil
}

//Update FOR UPDATING order information in database
func (db *SellerRepo) Update(id uint) (string, error) {
	message := "Update successfuly"
	return message, nil
}
