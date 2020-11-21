package app

import (
	"github.com/ahmed-bahaa/bookstore_users-api/controllers/ping"
	"github.com/ahmed-bahaa/bookstore_users-api/controllers/users"
)

func mapUrls() {

	// note we are putting the name of the function not calling it
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.PUT("/users:user_id", users.UpdateUser)
	router.GET("/users/:user_id", users.GetUser)

}
