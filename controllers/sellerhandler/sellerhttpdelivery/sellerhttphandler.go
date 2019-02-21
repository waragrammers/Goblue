package sellerhttpdelivery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/waragrammers/Goblue/controllers/sellerhandler/sellerrepo"
)

//SellerHTTPhandler interactor struct
type SellerHTTPhandler struct {
	sellerRepo sellerrepo.SellerRepo
}

//NewSellerHTTPhandler initiate order requests
func NewSellerHTTPhandler(route *echo.Echo) {
	handlefunc := &SellerHTTPhandler{sellerrepo.SellerRepo{}}
	order := route.Group("/seller")
	order.POST("/create", handlefunc.Create)
	order.GET("/single/:id", handlefunc.GetByID)
	order.GET("/all", handlefunc.FetchAll)
	order.PUT("/modify/:id", handlefunc.Update)
	order.DELETE("/delete/:id", handlefunc.Delete)
}

// Create for assigning shipper to order request
func (sellerH *SellerHTTPhandler) Create(c echo.Context) error {
	example, err := sellerH.sellerRepo.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//GetByID for get single order request
func (sellerH *SellerHTTPhandler) GetByID(c echo.Context) error {
	example, err := sellerH.sellerRepo.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//FetchAll for getall order request
func (sellerH *SellerHTTPhandler) FetchAll(c echo.Context) error {
	example, err := sellerH.sellerRepo.FechAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//Update for update request
func (sellerH *SellerHTTPhandler) Update(c echo.Context) error {
	example, err := sellerH.sellerRepo.Update(3)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//Delete handle delete request
func (sellerH *SellerHTTPhandler) Delete(c echo.Context) error {
	example, err := sellerH.sellerRepo.Delete(4)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}
