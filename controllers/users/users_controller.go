package users

import (
	"net/http"
	"strconv"

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
	result, Saveerr := services.CreateUser(user)
	if Saveerr != nil {
		c.JSON(Saveerr.Status, Saveerr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userID, UserErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if UserErr != nil {
		err := errors.NewBadRequestError("User ID should be a number")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userID)
	if UserErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, user)
}
