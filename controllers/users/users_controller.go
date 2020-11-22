package users

import (
	"net/http"
	"strconv"

	"github.com/ahmed-bahaa/bookstore_users-api/utils/errors"

	"github.com/ahmed-bahaa/bookstore_users-api/domain/users"
	"github.com/ahmed-bahaa/bookstore_users-api/services"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userID, UserErr := strconv.ParseInt(userIdParam, 10, 64)
	if UserErr != nil {
		err := errors.NewBadRequestError("User ID should be a number")
		return 0, err
	}
	return userID, nil
}

func Create(c *gin.Context) {

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

func Get(c *gin.Context) {

	userID, UserErr := getUserId(c.Param("user_id"))
	if UserErr != nil {
		c.JSON(UserErr.Status, UserErr)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {

	userID, UserErr := getUserId(c.Param("user_id"))
	if UserErr != nil {
		c.JSON(UserErr.Status, UserErr)
		return
	}

	var userUpdated users.User
	err := c.ShouldBindJSON(&userUpdated)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	userUpdated.ID = userID
	result, updateErr := services.UpdateUser(isPartial, userUpdated)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userID, UserErr := getUserId(c.Param("user_id"))
	if UserErr != nil {
		c.JSON(UserErr.Status, UserErr)
		return
	}
	err := services.DeleteUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"Status": "deleted"})
}
