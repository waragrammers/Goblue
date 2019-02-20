package orderdeliverly

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/waragrammers/Goblue/controllers/orderhandler/oderrepo"
)

//OrderHTTPhandler interactor struct
type OrderHTTPhandler struct {
	OrderRepo oderrepo.OrderRepo
}

//NewOrderHTTPhandler initiate order requests
func NewOrderHTTPhandler(route *echo.Echo) {
	handlefunc := &OrderHTTPhandler{oderrepo.OrderRepo{}}
	order := route.Group("/order")
	order.POST("/create", handlefunc.CreateOrder)
	order.GET("/order/:id", handlefunc.GetOrderByID)
	order.GET("/orders", handlefunc.GetOrders)
	order.PUT("/order/:id", handlefunc.UpdateOrder)
	order.DELETE("/order/:id", handlefunc.DeleteOrder)
}

//CreateOrder for saving order request
func (orderH *OrderHTTPhandler) CreateOrder(c echo.Context) error {
	example, err := orderH.OrderRepo.CreateOrder()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//GetOrderByID for get single order request
func (orderH *OrderHTTPhandler) GetOrderByID(c echo.Context) error {
	example, err := orderH.OrderRepo.GetOrderByID(2)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//GetOrders for getall order request
func (orderH *OrderHTTPhandler) GetOrders(c echo.Context) error {
	example, err := orderH.OrderRepo.GetOrders()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//UpdateOrder for update request
func (orderH *OrderHTTPhandler) UpdateOrder(c echo.Context) error {
	example, err := orderH.OrderRepo.UpdateOrder(3)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//DeleteOrder handle delete request
func (orderH *OrderHTTPhandler) DeleteOrder(c echo.Context) error {
	example, err := orderH.OrderRepo.DeleteOrder(4)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}
