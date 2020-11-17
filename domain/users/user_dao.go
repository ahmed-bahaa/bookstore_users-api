package users

import (
	"fmt"

	"github.com/ahmed-bahaa/bookstore_users-api/utils/date_utils"

	"github.com/ahmed-bahaa/bookstore_users-api/utils/errors"
)

var (
	userdb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	result := userdb[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
	}

	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {

	result := userdb[user.ID]
	if result != nil {
		if result.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exist", user.ID))
	}

	user.DateCreated = date_utils.GetNowString()
	userdb[user.ID] = user
	return nil
}
