package httpdelivery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/waragrammers/Goblue/controllers/locationhandler/locationrepo"
)

//LocationHTTPhandler interactor struct
type LocationHTTPhandler struct {
	LocationRepo locationrepo.LocationRepo
}

//NewOrderHTTPhandler initiate order requests
func NewOrderHTTPhandler(route *echo.Echo) {
	handlefunc := &LocationHTTPhandler{locationrepo.LocationRepo{}}
	order := route.Group("/address")
	order.POST("/location", handlefunc.SaveLocation)
	order.GET("/location/:id", handlefunc.GetByID)
	order.GET("/location", handlefunc.GetLocation)
	order.PUT("/location/:id", handlefunc.Update)
	order.DELETE("/location/:id", handlefunc.Delete)
}

//SaveLocation for saving  request
func (orderH *LocationHTTPhandler) SaveLocation(c echo.Context) error {
	example, err := orderH.LocationRepo.SaveLocation()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//GetByID for get single order request
func (orderH *LocationHTTPhandler) GetByID(c echo.Context) error {
	example, err := orderH.LocationRepo.GetByID(1)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//GetLocation for getall  request
func (orderH *LocationHTTPhandler) GetLocation(c echo.Context) error {
	example, err := orderH.LocationRepo.GetLocation()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//Update for update request
func (orderH *LocationHTTPhandler) Update(c echo.Context) error {
	example, err := orderH.LocationRepo.Update(2)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//Delete handle delete request
func (orderH *LocationHTTPhandler) Delete(c echo.Context) error {
	example, err := orderH.LocationRepo.Delete(3)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}
