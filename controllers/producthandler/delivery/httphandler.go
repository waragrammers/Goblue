package delivery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/waragrammers/Goblue/controllers/producthandler/repository"
)

//HTTPHandler will help us to handle request
type HTTPHandler struct {
	ProductRepo repository.ProductRepo
}

//NewProductHTTPHandler will be helping us to initiate new one
func NewProductHTTPHandler(route *echo.Echo) {
	handleFunc := &HTTPHandler{repository.ProductRepo{}}
	route.Use(middleware.Logger())
	route.GET("/allproducts", handleFunc.fechAll)
	route.POST("/create", handleFunc.CreateProduct)
	route.PUT("/update/:idpro", handleFunc.Update)
	route.DELETE("/delete/:idpro", handleFunc.Delete)
}

func (hndl *HTTPHandler) fechAll(c echo.Context) error {
	data, err := hndl.ProductRepo.FechAll()
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, data)
}

//CreateProduct handler used to create product
func (hndl *HTTPHandler) CreateProduct(c echo.Context) error {
	message := hndl.ProductRepo.Create()

	return c.String(http.StatusOK, message)
}

//Update handler used to update product
func (hndl *HTTPHandler) Update(c echo.Context) error {
	message := hndl.ProductRepo.Update(2)

	return c.String(http.StatusOK, message)
}

//Delete handler used to Delete product
func (hndl *HTTPHandler) Delete(c echo.Context) error {

	message := hndl.ProductRepo.Delete(2)

	return c.String(http.StatusOK, message)

}
