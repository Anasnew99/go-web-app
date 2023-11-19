package controllers

import (
	"anasnew99/server/chat_app/collections"
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/models"
	"anasnew99/server/chat_app/utils"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomController struct {
	GetRoomCollection     func() *mongo.Collection
	AddRoom               func(room models.AddRoomObject) (*mongo.InsertOneResult, error)
	JoinRoom              func(username string, roomId string, password string) error
	GetRoom               func(roomId string) (models.Room, error)
	GetMinimalizedRoom    func(roomId string, withPassword bool) (models.Room, error)
	IsUserJoinedInTheRoom func(username string, roomId string) (room models.Room, isJoined bool)
	LeaveRoom             func(username string, roomId string) error
	DeleteRoom            func(roomId string) error
	IsUserOwnerOfTheRoom  func(username string, roomId string) bool
	getUserJoinedRooms    func(username string) []models.Room
}

func getRoomCollection() *mongo.Collection {
	return db.GetDB().Collection(collections.ROOMS)
}

func addRoom(room models.AddRoomObject) (*mongo.InsertOneResult, error) {
	room.CreatedAt = time.Now().Unix()
	if room.Password != "" {
		room.Password = utils.GetHashedString(room.Password)

	}
	room.Users = []string{}
	go addUserRoom(room.RoomOwner, room.Id)
	return getRoomCollection().InsertOne(context.TODO(), room)
}

func joinRoom(username string, roomId string, password string) error {
	room, e := getMinimalizedRoom(roomId, true)
	if err := e; err != nil {
		return err
	}

	if room.Password != "" && room.Password != utils.GetHashedString(password) {
		return errors.New("wrong password")
	}

	_, err := getRoomCollection().UpdateOne(context.TODO(), bson.M{
		"_id": roomId,
	}, bson.M{
		"$addToSet": bson.M{
			"users": username,
		},
	})
	return err

}

func getRoom(roomId string) (models.Room, error) {
	var room models.Room
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: collections.USERS},
		{Key: "localField", Value: "room_owner"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "room_owner"},
	},
	}}

	lookupStage2 := bson.D{
		{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: collections.USERS},
				{Key: "localField", Value: "users"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "users"},
			},
		},
	}
	// use $unwind to flatten the room_owner array
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{
		{Key: "path", Value: "$room_owner"},
	}}}
	// use $match to filter by roomId
	matchStage := bson.D{{Key: "$match", Value: bson.D{
		{Key: "_id", Value: roomId},
	}}}
	projectStage := bson.D{{Key: "$project", Value: bson.D{
		{Key: "password", Value: 0},
		{Key: "room_owner.password", Value: 0},
		{Key: "room_owner.rooms", Value: 0},
		{Key: "users.password", Value: 0},
		{Key: "users.rooms", Value: 0},
		{Key: "messages", Value: 0},
	}}}
	pipeline := mongo.Pipeline{matchStage, lookupStage, lookupStage2, unwindStage, projectStage}

	cursor, err := getRoomCollection().Aggregate(context.Background(), pipeline)

	if err != nil {
		return room, err
	}
	defer cursor.Close(context.Background())
	var data any

	if cursor.Next(context.Background()) {
		err = cursor.Decode(&room)
		cursor.Decode(&data)
		if err != nil {
			return room, err
		}
	} else {

		return room, mongo.ErrNoDocuments
	}
	return room, nil
}

func getMinimalizedRoom(roomId string, withPassword bool) (models.Room, error) {
	var room models.Room
	// use $lookup to get the User object for RoomOwner
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: collections.USERS},
		{Key: "localField", Value: "room_owner"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "room_owner"},
	},
	}}
	// use $unwind to flatten the room_owner array
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{
		{Key: "path", Value: "$room_owner"},
	}}}
	// use $match to filter by roomId
	matchStage := bson.D{{Key: "$match", Value: bson.D{
		{Key: "_id", Value: roomId},
	}}}
	// use $project to exclude the password field
	projectStage := bson.D{{Key: "$project", Value: bson.D{
		{Key: "room_owner.password", Value: 0},
		{Key: "room_owner.rooms", Value: 0},
		{Key: "messages", Value: 0},
		{Key: "users", Value: 0},
	}}}
	// use $sort to sort messages by createdAt in descending order

	// use $limit to limit the number of messages to 50
	// aggregate the pipeline stages

	pipeline := mongo.Pipeline{matchStage, lookupStage, unwindStage, projectStage}
	cursor, err := getRoomCollection().Aggregate(context.Background(), pipeline)
	if err != nil {
		return room, err
	}
	defer cursor.Close(context.Background())
	if cursor.Next(context.Background()) {
		err = cursor.Decode(&room)
		if err != nil {
			return room, err
		}
	} else {
		return room, mongo.ErrNoDocuments
	}
	if !withPassword {
		room.Password = ""
	}
	return room, nil
}

func isUserJoinedInTheRoom(username string, roomId string) (room models.Room, isJoined bool) {
	var r models.Room
	r, err := getRoom(roomId)
	room = r
	if err != nil {
		return room, false
	}
	if r.RoomOwner.Username == username {
		return room, true
	}

	for _, user := range r.Users {
		if user.Username == username {
			return room, true
		}
	}
	return room, false

}

func leaveRoom(username string, roomId string) error {
	_, err := getRoomCollection().UpdateOne(context.TODO(), bson.M{
		"_id": roomId,
	}, bson.M{
		"$pull": bson.M{
			"users": username,
		},
	})
	return err
}

func deleteRoom(roomId string) error {

	_, err := getRoomCollection().DeleteOne(context.TODO(), bson.M{
		"_id": roomId,
	})
	return err
}

func isUserOwnerOfTheRoom(username string, roomId string) bool {
	room, err := getMinimalizedRoom(roomId, false)
	if err != nil {
		return false
	}
	return room.RoomOwner.Username == username
}

func getUserJoinedRooms(username string) []models.Room {
	var rooms []models.Room
	cursor, err := getRoomCollection().Find(context.TODO(), bson.M{
		"users": username,
	}, options.Find().SetProjection(bson.M{
		"password":   0,
		"users":      0,
		"room_owner": 0,
		"messages":   0,
	}))

	if err != nil {
		return rooms
	}

	defer cursor.Close(context.Background())
	cursor.All(context.Background(), &rooms)

	return rooms
}

var Room = &RoomController{
	GetRoomCollection:     getRoomCollection,
	AddRoom:               addRoom,
	JoinRoom:              joinRoom,
	GetRoom:               getRoom,
	GetMinimalizedRoom:    getMinimalizedRoom,
	IsUserJoinedInTheRoom: isUserJoinedInTheRoom,
	LeaveRoom:             leaveRoom,
	DeleteRoom:            deleteRoom,
	IsUserOwnerOfTheRoom:  isUserOwnerOfTheRoom,
	getUserJoinedRooms:    getUserJoinedRooms,
}
