package main

import (
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/waragrammers/Goblue/controllers/locationhandler/httpdelivery"
	"github.com/waragrammers/Goblue/controllers/locationhandler/locationrepo"
	"github.com/waragrammers/Goblue/controllers/orderhandler/oderrepo"
	"github.com/waragrammers/Goblue/controllers/orderhandler/orderdeliverly"
	"github.com/waragrammers/Goblue/controllers/producthandler/productdelivery"
	"github.com/waragrammers/Goblue/controllers/producthandler/repository"
	"github.com/waragrammers/Goblue/controllers/shipperorderhandler/httpdeliverly"
	"github.com/waragrammers/Goblue/controllers/shipperorderhandler/shipperorderrepo"
	"github.com/waragrammers/Goblue/controllers/userhandler/userdelivery"
	"github.com/waragrammers/Goblue/controllers/userhandler/userrepo"
	"github.com/waragrammers/Goblue/dbdriver"
)

func main() {

	DbConnection := dbdriver.PostgresDb()
	echoRoute := echo.New()
	//ProductHandler
	repository.NewProductRepo(DbConnection)
	productdelivery.NewProductHTTPHandler(echoRoute)
	//UserHandler
	userrepo.NewUserRepostory(DbConnection)
	userdelivery.NewUserHTTPhandler(echoRoute)

	//orderHandler
	oderrepo.NewOrderRepo(DbConnection)
	orderdeliverly.NewOrderHTTPhandler(echoRoute)

	//shipperorder handler
	shipperorderrepo.NewShipperOrderrepo(DbConnection)
	httpdeliverly.NewShipperOrderHTTPhandler(echoRoute)

	//Location handler
	locationrepo.NewLocationRepo(DbConnection)
	httpdelivery.NewOrderHTTPhandler(echoRoute)
	echoRoute.Start(":3000")
}
