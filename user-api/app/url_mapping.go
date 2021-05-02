package app

import "github.com/golang-microservice/user-api/controllers"

func mapUrls()  {
		router.GET("/ping",controllers.Ping)
		router.POST("/users",controllers.CreateUser)
		router.PUT("/users/:user_id",controllers.UpdateUser)
		router.PATCH("/users/:user_id",controllers.UpdateUser)
		router.GET("/users/:user_id",controllers.GetUser)
		
		// router.GET("/users/search",controllers.SearchUser)
		
}