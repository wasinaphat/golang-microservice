package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/golang-microservice/oauth-api/clients/cassandra"
	"github.com/golang-microservice/oauth-api/domain/access_token"
	"github.com/golang-microservice/oauth-api/http"
	"github.com/golang-microservice/oauth-api/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {

	fmt.Println("START APPLICATION")

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
