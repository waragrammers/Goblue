package locationrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/controllers/locationhandler"
	"github.com/waragrammers/Goblue/models"
)

//LocationRepo  for brouth the database in you arsenal
type LocationRepo struct {
	DbConnection *gorm.DB
}

// NewLocationRepo for initialize new LocationRepo object
func NewLocationRepo(DbCon *gorm.DB) locationhandler.LocationRepostory {
	return &LocationRepo{DbCon}
}

//SaveLocation for saving the order in database
func (db *LocationRepo) SaveLocation() (*models.LocationAdress, error) {
	//database access
	return nil, nil
}

//GetByID for get single locationInfo in database
func (db *LocationRepo) GetByID(id uint) (*models.LocationAdress, error) {

	//database access
	return nil, nil
}

//GetLocation get all locationsin database
func (db *LocationRepo) GetLocation() ([]*models.LocationAdress, error) {
	//database access
	return nil, nil
}

//Delete for delete single location record in database
func (db *LocationRepo) Delete(id uint) (*models.LocationAdress, error) {
	//database access
	return nil, nil
}

//Update FOR UPDATING location information in database
func (db *LocationRepo) Update(id uint) (string, error) {
	message := "Update successfuly"
	return message, nil
}
