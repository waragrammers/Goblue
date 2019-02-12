package models

import "database/sql"

//Product model will be used as product table
type Product struct {
	ID                 uint
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
	ProductImage       []*ProductImage
	Order              []*Order    `gorm:"many2many:product_order;"`
	Category           []*Category `gorm:"many2many:product_category;"`
}

//Order model will be used as order table
type Order struct {
	ID         uint
	UserID     uint
	Address    string
	Product    []*Product `gorm:"many2many:product_order;"`
	DistanceID uint
	ShipperID  uint
}

//Distance model will be used as Distance table
type Distance struct {
	ID      uint
	LongOne float64
	LatOne  float64
	LongTwo float64
	LatTwo  float64
	Order   []*Order
}

//Category model will be used as Category table
type Category struct {
	ID            uint
	CategoryTitle string     `gorm:"type:varchar(100);not null"`
	Product       []*Product `gorm:"many2many:product_category;"`
}

//Shipper model will be used as shiper table
type Shipper struct {
	ID           uint
	FirstName    string `gorm:"type:varchar(100);not null"`
	LastName     string `gorm:"type:varchar(100);not null"`
	Phone        string `gorm:"type:varchar(100);unique;not null"`
	email        string `gorm:"type:varchar(100);unique;not null"`
	ProfileImage string
	Order        []*Order
	User         User
	UserID       uint
}

//OrderDetails model will be used as orderdetails table
type OrderDetails struct {
	ID           uint
	IsDeliverded sql.NullBool
	Order        *Order
	OrderID      uint
}

//Seller model will be used as seller table
type Seller struct {
	ID            uint
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
	ID        uint
	ProductID uint
}

//User model will be used as user table
type User struct {
	ID       uint
	UserName string
	email    string `gorm:"type:varchar(100);unique;not null"`
	password string `gorm:"not null"`
}

//ProductImage model will be used as productimage table
type ProductImage struct {
	ID         uint
	ImageOne   sql.NullString
	ImageTwo   sql.NullString
	ImageThree sql.NullString
	ProductID  uint
}
