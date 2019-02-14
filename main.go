package main

import (
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/waragrammers/Goblue/controllers/producthandler/delivery"
	"github.com/waragrammers/Goblue/controllers/producthandler/repository"
	"github.com/waragrammers/Goblue/dbdriver"
)

func main() {

	DbConnection := dbdriver.PostgresDb()
	echoRoute := echo.New()

	repository.NewProductRepo(DbConnection)
	delivery.NewProductHTTPHandler(echoRoute)
	echoRoute.Start(":3000")
}
