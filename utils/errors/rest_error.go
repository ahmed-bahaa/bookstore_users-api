package errors

import (
	"net/http"
)

type RestErr struct {
	Message string `JSON:"message"`
	Status  int    `JSON:"status"`
	Error   string `JSON:"error"`
}

func NewBadRequestError(messgae string) *RestErr {
	return &RestErr{
		Message: messgae,
		Status:  http.StatusBadRequest,
		Error:   "Bad_Request",
	}
}

func NewNotFoundError(messgae string) *RestErr {
	return &RestErr{
		Message: messgae,
		Status:  http.StatusNotFound,
		Error:   "Not_found",
	}
}

func NewInteralServerError(messgae string) *RestErr {
	return &RestErr{
		Message: messgae,
		Status:  http.StatusInternalServerError,
		Error:   "Not_found",
	}
}

// errors to be handeled

// {
// 	"message": "user 123 not found"
// 	"status": 404
// 	"error": "Not_found"
// }

// {
// 	"message": "Invalid json body"
// 	"status": 400
// 	"error": "Bad_request"
// }

// {
// 	"message": "database is down"
// 	"status": 500
// 	"error": "Internal_server_error"
// }
