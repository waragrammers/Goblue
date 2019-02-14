package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/controllers/producthandler"
	"github.com/waragrammers/Goblue/models"
)

// ProductRepo will be used to implement GeneralProductRepo method
type ProductRepo struct {
	DbConn *gorm.DB
}

// NewProductRepo will be used to return db
func NewProductRepo(DbConnection *gorm.DB) producthandler.GeneralProductRepo {
	return &ProductRepo{DbConnection}
}

//FechAll will return all product
func (db *ProductRepo) FechAll() ([]*models.Product, error) {

	return nil, nil
}

//Create will deal with creating new product
func (db *ProductRepo) Create() (message string) {
	create := "This is create function"
	return create
}

//Update just to update product based on its ID
func (db *ProductRepo) Update(id int) (message string) {
	update := "this is update"
	return update
}

//Delete product based on its ID
func (db *ProductRepo) Delete(id int) (message string) {
	delete := "this is delelte methode"
	return delete
}
