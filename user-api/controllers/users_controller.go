package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-microservice/user-api/domain/users"
	"github.com/golang-microservice/user-api/services"
	"github.com/golang-microservice/user-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		// TODO : return bad request to the caller
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		// TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	result,getErr := services.GetUser(userId)
	if getErr!=nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}



func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Search User")
}
