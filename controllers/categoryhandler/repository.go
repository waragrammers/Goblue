package categoryhandler

import (
	"github.com/waragrammers/Goblue/models"
)

//CategoryRepository  exposes all method to will be implemented
type CategoryRepository interface {
	Create() (*models.Category, error)
	FechAll() ([]*models.Category, error)
	GetByID(id uint) (*models.Category, error)
	Delete(id uint) (*models.Category, error)
	Update(id uint) (string, error)
}
