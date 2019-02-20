package locationhandler

import (
	"github.com/waragrammers/Goblue/models"
)

//LocationRepostory method operation for order
type LocationRepostory interface {
	SaveLocation() (*models.LocationAdress, error)
	GetLocation() ([]*models.LocationAdress, error)
	GetByID(id uint) (*models.LocationAdress, error)
	Delete(id uint) (*models.LocationAdress, error)
	Update(id uint) (string, error)
}
