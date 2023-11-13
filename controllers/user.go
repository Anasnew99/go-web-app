package controllers

import (
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

func GetUserCollection() *mongo.Collection {
	return db.GetDB().Collection("users")
}

func InsertUser(user User) (*mongo.InsertOneResult, error) {
	user.Password = utils.GetHashedString(user.Password)
	user.Id = user.Username

	return GetUserCollection().InsertOne(context.TODO(), user)
}

func DeleteUser(username string) (*mongo.DeleteResult, error) {
	return GetUserCollection().DeleteOne(context.TODO(), bson.M{
		"_id": username,
	})
}

func GetUser(username string) (User, error) {
	var user User
	err := GetUserCollection().FindOne(context.TODO(), bson.M{
		"_id": username,
	}).Decode(&user)
	return user, err
}
