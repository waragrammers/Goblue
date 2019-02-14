package delivery

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
	route.POST("/user/create", handleFunc.CreateUser)
}

//CreateUser for creating new user
func (userH *UserHTTPhandler) CreateUser(c echo.Context) error {
	message, error := userH.UserRepo.CreateUser()
	if error != nil {
		return nil
	}

	return c.String(http.StatusOK, message)

}

// CreateUser() (string, error)
// 	GetAllUser() ([]*models.User, error)
// 	GetuserByID(id int) (*models.User, error)
// 	UpdateUser(id int) (string, error)
