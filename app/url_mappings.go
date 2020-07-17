package app

import (
	"github.com/sumudhar/go-book-store-user-api/controllers/ping"
	"github.com/sumudhar/go-book-store-user-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping",ping.Ping)

	router.POST("/users",users.Create)
	router.GET("/users/:user_id",users.Get)
	router.PUT("/users/:user_id",users.Update)
	router.PATCH("/users/:user_id",users.Update)
	router.DELETE("/users/:user_id",users.Delete)
	//router.GET("/users/search",users.SearchUser)
	

}