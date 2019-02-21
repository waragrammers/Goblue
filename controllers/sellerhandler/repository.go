package sellerhandler

import (
	"github.com/waragrammers/Goblue/models"
)

//SellerRepository exposing method will be implemented
type SellerRepository interface {
	Create() (*models.Seller, error)
	FechAll() ([]*models.Seller, error)
	GetByID(id uint) (*models.Seller, error)
	Delete(id uint) (*models.Seller, error)
	Update(id uint) (string, error)
}
