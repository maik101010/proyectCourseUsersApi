package app

import (
	"github.com/maik101010/proyectCourseUsersApi/controller/ping"
	user "github.com/maik101010/proyectCourseUsersApi/controller/users"
)

//MapUrls method
func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", user.Create)
	router.GET("/users/:user_id", user.Get)
	router.PUT("/users/:user_id", user.Update)
	router.PATCH("/users/:user_id", user.Update)
	router.DELETE("/users/:user_id", user.Delete)
	router.GET("/internal/users/search", user.Search)
	router.POST("/users/login", user.Login)
	//Truncate user
	router.GET("/internal/users/truncate", user.Truncate)
	// router.GET("/users/search", controller.SearchUser)

}
