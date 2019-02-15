package main

import (
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/waragrammers/Goblue/controllers/producthandler/productdelivery"
	"github.com/waragrammers/Goblue/controllers/producthandler/repository"
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
	echoRoute.Start(":3000")
}
