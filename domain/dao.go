package domain

import (
	"context"
	"time"

	"github.com/parsaakbari1209/go-mongo-CRUD-web/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert user to the database.")
		return nil, restErr
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}
