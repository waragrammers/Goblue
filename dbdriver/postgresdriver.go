package dbdriver

import (
	"github.com/jinzhu/gorm"
	"github.com/waragrammers/Goblue/models"
)

//PostgresDb should return database to any module want to use it
func PostgresDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5500 user=postgres dbname=postgres password=gocore sslmode=disable")
	defer db.Close()
	if err != nil {
		panic("Connection Fail")
	}
	db.SingularTable(true)

	db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.Distance{},
		&models.Order{},
		&models.OrderDetails{},
		&models.ProductImage{},
		&models.Ranking{},
		&models.Seller{},
		&models.Shipper{},
		&models.User{})

	return db
}
