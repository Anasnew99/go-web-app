package controllers

import (
	"anasnew99/server/chat_app/collections"
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/models"
	"anasnew99/server/chat_app/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserController struct {
	GetUserCollection  func() *mongo.Collection
	AddUser            func(user models.User) (*mongo.InsertOneResult, error)
	DeleteUser         func(username string) (*mongo.DeleteResult, error)
	GetUserRooms       func(username string) ([]models.Room, error)
	GetUser            func(username string) (models.User, error)
	ChangeUserPassword func(username string, oldPassword string, newPassword string) error
	AddUserRoom        func(username string, roomId string) error
}

func getUserCollection() *mongo.Collection {
	return db.GetDB().Collection(collections.USERS)
}

func addUser(user models.User) (*mongo.InsertOneResult, error) {
	user.Password = utils.GetHashedString(user.Password)

	user.Id = user.Username
	user.Rooms = []models.Room{}

	return getUserCollection().InsertOne(context.TODO(), user)
}

func deleteUser(username string) (*mongo.DeleteResult, error) {
	return getUserCollection().DeleteOne(context.TODO(), bson.M{
		"_id": username,
	})
}

func getUserRooms(username string) ([]models.Room, error) {
	// get rooms object from room collection
	var data struct {
		Rooms []string `json:"rooms" bson:"rooms"`
	}
	err := getUserCollection().FindOne(context.TODO(), bson.M{
		"_id": username,
	}, options.FindOne().SetProjection(bson.M{
		"rooms": 1,
	})).Decode(&data)

	if err != nil {
		return []models.Room{}, err
	}
	var rooms []models.Room
	for _, roomId := range data.Rooms {
		var room models.Room
		room, err := Room.GetMinimalizedRoom(roomId, false)
		if err == nil {
			rooms = append(rooms, room)
		}
	}

	joinedRooms := Room.getUserJoinedRooms(username)
	rooms = append(rooms, joinedRooms...)
	return rooms, nil
}

func getUser(username string) (models.User, error) {
	var user models.User
	// lookup rooms
	err := getUserCollection().FindOne(context.TODO(), bson.M{
		"_id": username,
	}, options.FindOne().SetProjection(bson.M{
		"rooms": 0,
	})).Decode(&user)
	user.Rooms, _ = getUserRooms(username)
	return user, err
}

func changeUserPassword(username string, oldPassword string, newPassword string) error {

	if user, err := getUser(username); err != nil || user.Password != utils.GetHashedString(oldPassword) {
		return errors.New("wrong password")
	}
	getUserCollection().UpdateOne(context.TODO(), bson.M{
		"_id": username,
	}, bson.M{
		"$set": bson.M{
			"password": utils.GetHashedString(newPassword),
		},
	})
	return nil
}

func addUserRoom(username string, roomId string) error {
	_, err := getUserCollection().UpdateOne(context.TODO(), bson.M{
		"_id": username,
	}, bson.M{
		"$addToSet": bson.M{
			"rooms": roomId,
		},
	})
	return err
}

var User = &UserController{
	GetUserCollection:  getUserCollection,
	AddUser:            addUser,
	DeleteUser:         deleteUser,
	GetUserRooms:       getUserRooms,
	GetUser:            getUser,
	ChangeUserPassword: changeUserPassword,
	AddUserRoom:        addUserRoom,
}
