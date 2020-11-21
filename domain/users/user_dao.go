package users

import (
	"github.com/ahmed-bahaa/bookstore_users-api/utils/mysql_utils"

	"github.com/ahmed-bahaa/bookstore_users-api/utils/date_utils"

	"github.com/ahmed-bahaa/bookstore_users-api/datasources/mysql/useres_db"

	"github.com/ahmed-bahaa/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
	indexUniqueEmail = "email_UNIQUE"
	// errorNoRows      = "no rows in result set"
)

func (user *User) Get() *errors.RestErr {

	stmt, err := useres_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInteralServerError(err.Error())
	}
	// defer ,, it is a stack with functions which will run before exectuing the return statment
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)

	err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := useres_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInteralServerError(err.Error())
	}
	// defer ,, it is a stack with functions which will run before exectuing the return statment
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertionResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userID, err := insertionResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.ID = userID

	// if result != nil {
	// 	if result.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("User %d already exist", user.ID))
	// }

	// user.DateCreated = date_utils.GetNowString()
	// userdb[user.ID] = user
	return nil
}
