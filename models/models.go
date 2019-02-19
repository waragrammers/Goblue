package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

//Product model will be used as product table
type Product struct {
	gorm.Model
	ProductName        string `gorm:"type:varchar(100);not null"`
	ProductDescription string `gorm:"type:varchar(255);not null"`
	QuantityPerUnit    float32
	UnitPrice          float64 `gorm:"not null"`
	AvailablePrice     float64
	AvailableColor     sql.NullString
	Size               sql.NullString
	Color              sql.NullString
	Discount           float64
	UnitWeight         sql.NullString
	UnitsInStock       float64
	UnitsOnOrder       string
	ProductAvailable   string
	SellerID           uint
	Ranking            []*Ranking
	ProductImage       *ProductImage
	Order              []*Order `gorm:"many2many:product_order;"`
	CategoryID         uint
}

//Order model will be used as order table
type Order struct {
	gorm.Model
	UserID         uint
	Address        string
	Product        []*Product `gorm:"many2many:product_order;"`
	LocationAdress []*LocationAdress
}

//OrderDetails model will be used as orderdetails table
type OrderDetails struct {
	gorm.Model
	IsDeliverded sql.NullBool
	Order        *Order
	OrderID      uint
}

//ShipperOder model will be used as ShipperOder table
type ShipperOder struct {
	gorm.Model
	ShipperID uint
	Order     *Order
}

//LocationAdress  model will be used as LocationAdress  table
type LocationAdress struct {
	gorm.Model
	OrderID      uint
	FullName     string `gorm:"type:varchar(100);not null"`
	AddressLine1 string
	AddressLine2 string
	City         string
	Province     string
	Country      string
	PhoneNumber  string `gorm:"type:varchar(100);not null"`
}

//Category model will be used as Category table
type Category struct {
	gorm.Model
	CategoryTitle string `gorm:"type:varchar(100);not null"`
	Product       []*Product
}

//Shipper model will be used as shiper table
type Shipper struct {
	gorm.Model
	FirstName    string `gorm:"type:varchar(100);not null"`
	LastName     string `gorm:"type:varchar(100);not null"`
	Phone        string `gorm:"type:varchar(100);unique;not null"`
	email        string `gorm:"type:varchar(100);unique;not null"`
	ProfileImage string
	User         User
	UserID       uint
	ShipperOder  []*ShipperOder
}

//Seller model will be used as seller table
type Seller struct {
	gorm.Model
	CompanyName   string `gorm:"type:varchar(100);unique;not null"`
	Address1      string
	Address2      string
	City          string
	PostCode      sql.NullString
	Phone         string `gorm:"type:varchar(100);unique;not null"`
	Email         string `gorm:"type:varchar(100);unique;not null"`
	WebURL        sql.NullString
	DiscountAvail sql.NullFloat64
	PinNumber     string `gorm:"type:varchar(100);unique;not null"`
	Logo          string
	Product       []*Product
	User          User
	UserID        uint
}

//Ranking model will be used as ranking table
type Ranking struct {
	gorm.Model
	ProductID uint
}

//User model will be used as user table
type User struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar(100);unique;not null"`
	LastName    string `gorm:"type:varchar(100);unique;not null"`
	PhoneNumber string `gorm:"type:varchar(100);unique;not null"`
	email       string `gorm:"type:varchar(100);unique;not null"`
	password    string `gorm:"not null"`

	Role  []*Role `gorm:"many2many:user_role;"`
	Order []*Order
}

//ProductImage model will be used as productimage table
type ProductImage struct {
	gorm.Model
	ImageOne   sql.NullString
	ImageTwo   sql.NullString
	ImageThree sql.NullString
	ProductID  uint
}

//Role model will be used as Role table
type Role struct {
	gorm.Model
	RoleName string
	User     []*User `gorm:"many2many:user_role;"`
}
