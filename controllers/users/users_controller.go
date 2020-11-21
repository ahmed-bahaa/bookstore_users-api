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
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userID, UserErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if UserErr != nil {
		err := errors.NewBadRequestError("User ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	var userUpdated = users.User
	err := c.ShouldBindJSON(&userUpdated)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	userUpdated.ID = userID
	result, updateErr := services.UpdateUser(userUpdated)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, userUpdated)
}
