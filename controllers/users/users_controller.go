package users

import (
	"net/http"

	"github.com/ahmed-bahaa/bookstore_users-api/utils/errors"

	"github.com/ahmed-bahaa/bookstore_users-api/domain/users"
	"github.com/ahmed-bahaa/bookstore_users-api/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		resterr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(resterr.Status, resterr)
		return
	}
	// ==========> another aproach
	// fmt.Println("user:", user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// handeling the error
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(bytes))
	// err = json.Unmarshal(bytes, &user)
	// if err != nil {
	// 	// handeling the error
	// 	fmt.Println(err)
	// 	return
	// }

	result, Saveerr := services.CreateUser(user)
	if Saveerr != nil {
		c.JSON(Saveerr.Status, Saveerr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
