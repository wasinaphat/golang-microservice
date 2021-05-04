package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-microservice/oauth-api/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}

}


func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)

}
