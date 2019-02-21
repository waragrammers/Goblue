package categoryrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/controllers/categoryhandler"
	"github.com/waragrammers/Goblue/models"
)

//Categoryrepo for brouth the database in you arsenal
type Categoryrepo struct {
	DbConnection *gorm.DB
}

// NewCategoryrepo for initialize new orderrepo object
func NewCategoryrepo(DbCon *gorm.DB) categoryhandler.CategoryRepository {
	return &Categoryrepo{DbCon}
}

//Create for assign order in to shipper
func (db *Categoryrepo) Create() (*models.Category, error) {
	//database access
	return nil, nil
}

//GetByID for get single orderInfo in database
func (db *Categoryrepo) GetByID(id uint) (*models.Category, error) {

	//database access
	return nil, nil
}

//FechAll get all ordered orders in database
func (db *Categoryrepo) FechAll() ([]*models.Category, error) {
	//database access
	return nil, nil
}

//Delete for delete single order record in database
func (db *Categoryrepo) Delete(id uint) (*models.Category, error) {
	//database access
	return nil, nil
}

//Update FOR UPDATING order information in database
func (db *Categoryrepo) Update(id uint) (string, error) {
	message := "Update successfuly"
	return message, nil
}
