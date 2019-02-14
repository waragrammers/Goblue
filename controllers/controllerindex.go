package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/models"
)

type ProductRepo struct {
	DbConn *gorm.DB
}

func NewProductRepostory(DbConnection *gorm.DB) ProductRepo {
	return &ProductRepo{DbConnection}
}

func (db *ProductRepo) FechAll() ([]*models.Product, error) {

}
