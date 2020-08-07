package services

import (
	"github.com/parsaakbari1209/go-mongo-CRUD-web/domain"
	"github.com/parsaakbari1209/go-mongo-CRUD-web/utils"
)

func CreateUser(user *domain.User) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Create(user)
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}
