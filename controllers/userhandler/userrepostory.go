package userhandler

import (
	"github.com/waragrammers/Goblue/models"
)

//UserRepostory will be implemented by user handlers
type UserRepostory interface {
	CreateUser() (string, error)
	GetAllUser() ([]*models.User, error)
	GetuserByID(id int) (*models.User, error)
	UpdateUser(id int) (string, error)
	DeleteUser(id int) (string, error)
}
