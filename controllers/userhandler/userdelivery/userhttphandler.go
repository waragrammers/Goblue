package userdelivery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/waragrammers/Goblue/controllers/userhandler/userrepo"
)

//UserHTTPhandler for using to expose user URL
type UserHTTPhandler struct {
	UserRepo userrepo.UserRepo
}

//NewUserHTTPhandler for initializing this handler
func NewUserHTTPhandler(route *echo.Echo) {
	handleFunc := &UserHTTPhandler{userrepo.UserRepo{}}
	user := route.Group("/user")
	user.POST("/create", handleFunc.CreateUser)
	user.GET("/users", handleFunc.GetAllUser)
	user.GET("/user", handleFunc.GetuserByID)
	user.PUT("/update", handleFunc.UpdateUser)
	user.DELETE("/delete", handleFunc.DeleteUser)
}

//CreateUser handler for POST request
func (userH *UserHTTPhandler) CreateUser(c echo.Context) error {
	message, err := userH.UserRepo.CreateUser()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, message)

}

//GetAllUser handler for get request
func (userH *UserHTTPhandler) GetAllUser(c echo.Context) error {
	data, err := userH.UserRepo.GetAllUser()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

// GetuserByID handler for get/id request
func (userH *UserHTTPhandler) GetuserByID(c echo.Context) error {

	data, err := userH.UserRepo.GetuserByID(2)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

//UpdateUser handler for PUT request
func (userH *UserHTTPhandler) UpdateUser(c echo.Context) error {

	data, err := userH.UserRepo.UpdateUser(2)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, data)
}

//DeleteUser handler for delete request
func (userH *UserHTTPhandler) DeleteUser(c echo.Context) error {
	data, err := userH.UserRepo.DeleteUser(4)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, data)
}
