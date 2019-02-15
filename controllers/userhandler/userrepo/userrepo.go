package userrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/controllers/userhandler"
	"github.com/waragrammers/Goblue/models"
)

//UserRepo used to interact with BD
type UserRepo struct {
	DbConn *gorm.DB
}

//NewUserRepostory for initializing new userrepo
func NewUserRepostory(DbConnection *gorm.DB) userhandler.UserRepostory {
	return &UserRepo{DbConnection}
}

//CreateUser for creating user
func (db *UserRepo) CreateUser() (string, error) {
	message := "CreateUser function"
	return message, nil
}

//GetAllUser for getting all users
func (db *UserRepo) GetAllUser() ([]*models.User, error) {
	return nil, nil
}

//GetuserByID get single user
func (db *UserRepo) GetuserByID(id int) (*models.User, error) {
	return nil, nil
}

//UpdateUser update user info
func (db *UserRepo) UpdateUser(id int) (string, error) {
	message := "Updated"
	return message, nil
}

//DeleteUser method for deleting user in db
func (db *UserRepo) DeleteUser(id int) (string, error) {
	mesage := "Delete user method"

	return mesage, nil
}
