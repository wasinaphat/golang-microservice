package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"Create User")
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"Get User")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"Search User")
}
