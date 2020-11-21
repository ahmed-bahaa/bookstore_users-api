package services

import (
	"github.com/ahmed-bahaa/bookstore_users-api/domain/users"
	"github.com/ahmed-bahaa/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userId}
	getError := user.Get()
	if getError != nil {
		return nil, getError
	}
	return user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	err = user.Save()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	err = currentUser.Validate()
	if err != nil {
		return nil, err
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email

	err = currentUser.Update()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
