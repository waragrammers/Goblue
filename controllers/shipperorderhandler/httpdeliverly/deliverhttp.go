package httpdeliverly

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/waragrammers/Goblue/controllers/shipperorderhandler/shipperorderrepo"
)

//ShipperOrderHTTPhandler interactor struct
type ShipperOrderHTTPhandler struct {
	shipperorderrepo shipperorderrepo.ShipperOrderrepo
}

//NewShipperOrderHTTPhandler initiate order requests
func NewShipperOrderHTTPhandler(route *echo.Echo) {
	handlefunc := &ShipperOrderHTTPhandler{shipperorderrepo.ShipperOrderrepo{}}
	order := route.Group("/assign")
	order.POST("/shipper", handlefunc.AssignShipperToOrder)
	order.GET("/shipper/:id", handlefunc.GetByID)
	order.GET("/shipper", handlefunc.GetOrders)
	order.PUT("/shipper/:id", handlefunc.UpdateOrder)
	order.DELETE("/shipper/:id", handlefunc.DeleteOrder)
}

//AssignShipperToOrder for assigning shipper to order request
func (orderH *ShipperOrderHTTPhandler) AssignShipperToOrder(c echo.Context) error {
	example, err := orderH.shipperorderrepo.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//GetByID for get single order request
func (orderH *ShipperOrderHTTPhandler) GetByID(c echo.Context) error {
	example, err := orderH.shipperorderrepo.GetByID(2)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//GetOrders for getall order request
func (orderH *ShipperOrderHTTPhandler) GetOrders(c echo.Context) error {
	example, err := orderH.shipperorderrepo.GetOrders()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//UpdateOrder for update request
func (orderH *ShipperOrderHTTPhandler) UpdateOrder(c echo.Context) error {
	example, err := orderH.shipperorderrepo.UpdateOrder(3)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//DeleteOrder handle delete request
func (orderH *ShipperOrderHTTPhandler) DeleteOrder(c echo.Context) error {
	example, err := orderH.shipperorderrepo.DeleteOrder(4)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}
