package categorydelivery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/waragrammers/Goblue/controllers/categoryhandler/categoryrepo"
)

//CategoryHTTPhandler interactor struct
type CategoryHTTPhandler struct {
	categoryrepo categoryrepo.Categoryrepo
}

//NewCategoryHTTPhandler initiate order requests
func NewCategoryHTTPhandler(route *echo.Echo) {
	handlefunc := &CategoryHTTPhandler{categoryrepo.Categoryrepo{}}
	order := route.Group("/category")
	order.POST("/create", handlefunc.Create)
	order.GET("/single/:id", handlefunc.GetByID)
	order.GET("/all", handlefunc.FetchAll)
	order.PUT("/modify/:id", handlefunc.Update)
	order.DELETE("/delete/:id", handlefunc.Delete)
}

// Create for assigning shipper to order request
func (categoryH *CategoryHTTPhandler) Create(c echo.Context) error {
	example, err := categoryH.categoryrepo.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//GetByID for get single order request
func (categoryH *CategoryHTTPhandler) GetByID(c echo.Context) error {
	example, err := categoryH.categoryrepo.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//FetchAll for getall order request
func (categoryH *CategoryHTTPhandler) FetchAll(c echo.Context) error {
	example, err := categoryH.categoryrepo.FechAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//Update for update request
func (categoryH *CategoryHTTPhandler) Update(c echo.Context) error {
	example, err := categoryH.categoryrepo.Update(3)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}

//Delete handle delete request
func (categoryH *CategoryHTTPhandler) Delete(c echo.Context) error {
	example, err := categoryH.categoryrepo.Delete(4)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, example)
}
