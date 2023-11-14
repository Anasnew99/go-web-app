package controllers

import (
	"anasnew99/server/chat_app/collections"
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/utils"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
	Rooms    []Room `json:"rooms" bson:"rooms"`
}

func GetUserCollection() *mongo.Collection {
	return db.GetDB().Collection(collections.USERS)
}

func InsertUser(user User) (*mongo.InsertOneResult, error) {
	user.Password = utils.GetHashedString(user.Password)

	user.Id = user.Username
	user.Rooms = []Room{}

	return GetUserCollection().InsertOne(context.TODO(), user)
}

func DeleteUser(username string) (*mongo.DeleteResult, error) {
	return GetUserCollection().DeleteOne(context.TODO(), bson.M{
		"_id": username,
	})
}

func GetUserRooms(username string) ([]Room, error) {
	// get rooms object from room collection
	var data struct {
		Rooms []string `json:"rooms" bson:"rooms"`
	}
	err := GetUserCollection().FindOne(context.TODO(), bson.M{
		"_id": username,
	}, options.FindOne().SetProjection(bson.M{
		"rooms": 1,
	})).Decode(&data)

	if err != nil {
		return []Room{}, err
	}
	var rooms []Room
	for _, roomId := range data.Rooms {
		var room Room
		room, err := GetMinimalizedRoom(roomId)
		if err == nil {
			rooms = append(rooms, room)
		}
	}

	return rooms, nil
}

func GetUser(username string) (User, error) {
	var user User
	// lookup rooms
	err := GetUserCollection().FindOne(context.TODO(), bson.M{
		"_id": username,
	}, options.FindOne().SetProjection(bson.M{
		"rooms": 0,
	})).Decode(&user)
	user.Rooms, _ = GetUserRooms(username)
	return user, err
}

func ChangeUserPassword(username string, oldPassword string, newPassword string) error {

	if user, err := GetUser(username); err != nil || user.Password != utils.GetHashedString(oldPassword) {
		return errors.New("wrong password")
	}
	GetUserCollection().UpdateOne(context.TODO(), bson.M{
		"_id": username,
	}, bson.M{
		"$set": bson.M{
			"password": utils.GetHashedString(newPassword),
		},
	})
	return nil
}

func AddUserRoom(username string, roomId string) error {
	_, err := GetUserCollection().UpdateOne(context.TODO(), bson.M{
		"_id": username,
	}, bson.M{
		"$addToSet": bson.M{
			"rooms": roomId,
		},
	})
	fmt.Println(err)
	return err
}
