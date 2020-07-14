package app

import (
	"github.com/sumudhar/go-book-store-user-api/controllers/ping"
	"github.com/sumudhar/go-book-store-user-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping",ping.Ping)

	router.GET("/users/:user_id",users.GetUser)
	//router.GET("/users/search",users.SearchUser)
	router.POST("/users",users.CreateUser)
	

}