package mysql_utils

import (
	"strings"

	"github.com/ahmed-bahaa/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {

	mysqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(" No record matching the given id")
		}
		return errors.NewInteralServerError("Error parsing database response ")
	}

	switch mysqlError.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInteralServerError(" error,, Couldn't process your request")
}
