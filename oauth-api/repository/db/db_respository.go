package db

import (
	"fmt"

	"github.com/golang-microservice/oauth-api/clients/cassandra"
	"github.com/golang-microservice/oauth-api/domain/access_token"
	"github.com/golang-microservice/oauth-api/utils/errors"
)

const (
	queryGetAccessToken = "SELECT access_token,user_id,client_id, expires FROM access_tokens WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}
type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	//TODO : implement get access token from CassandraDB
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	fmt.Println("ID", id)
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil

}
