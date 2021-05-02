package app

import "github.com/golang-microservice/user-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.Create)
	router.PUT("/users/:user_id", controllers.Update)
	router.PATCH("/users/:user_id", controllers.Update)
	router.GET("/users/:user_id", controllers.Get)
	router.DELETE("/users/:user_id", controllers.Delete)

	// router.GET("/users/search",controllers.SearchUser)

}
