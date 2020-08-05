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

func Find(email string) (*User, *utils.RestErr) {
	var user User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := usersC.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	return &user, nil
}

func Delete(email string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		restErr := utils.NotFound("faild to delete.")
		return restErr
	}
	if result.DeletedCount == 0 {
		restErr := utils.NotFound("user not found.")
		return restErr
	}
	return nil
}

func Update(email string, field string, value string) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{field: value}})
	if err != nil {
		restErr := utils.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.BadRequest("no such field")
		return nil, restErr
	}
	user, restErr := Find(email)
	if restErr != nil {
		return nil, restErr
	}
	return user, restErr
}
