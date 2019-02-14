package producthandler

import "github.com/waragrammers/Goblue/models"

// GeneralProductRepo that will be use to handle all functionality applied to product table
type GeneralProductRepo interface {
	FechAll() ([]*models.Product, error)
	Create() (message string)
	Update(id int) (message string)
	Delete(id int) (message string)
}
