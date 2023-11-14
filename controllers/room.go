package controllers

import (
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Room struct {
	Id          string    `json:"id" bson:"_id"`
	Description string    `json:"description" bson:"description"`
	Password    string    `json:"password" bson:"password"`
	RoomOwner   User      `json:"room_owner" bson:"room_owner"`
	CreatedAt   int64     `json:"created_at" bson:"created_at"`
	Users       []User    `json:"users" bson:"users"`
	Messages    []Message `json:"messages" bson:"messages"`
}

type AddRoomObject struct {
	Description string `json:"description" bson:"description"`
	Password    string `json:"password" bson:"password"`
	RoomOwner   string `json:"room_owner" bson:"room_owner"`
	CreatedAt   int64  `json:"created_at" bson:"created_at"`
	Id          string `json:"id" bson:"_id"`
}

const ROOMS = "rooms"

func GetRoomCollection() *mongo.Collection {
	return db.GetDB().Collection(ROOMS)
}

func AddRoom(room AddRoomObject) (*mongo.InsertOneResult, error) {
	room.CreatedAt = time.Now().Unix()
	if room.Password != "" {
		room.Password = utils.GetHashedString(room.Password)

	}

	go AddUserRoom(room.RoomOwner, room.Id)
	return GetRoomCollection().InsertOne(context.TODO(), room)
}

func JoinRoom(username string, roomId string, password string) error {
	room, e := GetMinimalizedRoom(roomId)
	if err := e; err != nil {
		return err
	}
	if room.Password != "" && room.Password != utils.GetHashedString(password) {
		return errors.New("wrong password")
	}

	_, err := GetRoomCollection().UpdateOne(context.TODO(), bson.M{
		"_id": roomId,
	}, bson.M{
		"$addToSet": bson.M{
			"users": username,
		},
	})
	return err

}

func GetRoom(roomId string) (Room, error) {
	var room Room
	var data any
	// use $lookup to get the User object for RoomOwner
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: USERS},
		{Key: "localField", Value: "room_owner"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "room_owner"},
	},
	}}

	lookupStage2 := bson.D{
		{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: USERS},
				{Key: "let", Value: bson.D{{Key: "userIds", Value: "$users"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$in", Value: bson.A{"$_id", "$$userIds"}}}}}}},
				}},
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
	// use $project to exclude the password field
	projectStage := bson.D{{Key: "$project", Value: bson.D{
		{Key: "password", Value: 0},
		{Key: "room_owner.password", Value: 0},
		{Key: "room_owner.rooms", Value: 0},
		{Key: "users.password", Value: 0},
		{Key: "messages", Value: 0},
		{Key: "users.rooms", Value: 0},
	}}}
	// use $sort to sort messages by createdAt in descending order

	// use $limit to limit the number of messages to 50
	// aggregate the pipeline stages

	pipeline := mongo.Pipeline{matchStage, lookupStage, lookupStage2, unwindStage, projectStage}
	cursor, err := GetRoomCollection().Aggregate(context.Background(), pipeline)
	if err != nil {
		return room, err
	}
	defer cursor.Close(context.Background())
	if cursor.Next(context.Background()) {
		cursor.Decode(&data)
		err = cursor.Decode(&room)

		fmt.Println(data)
		if err != nil {
			return room, err
		}
	} else {
		return room, mongo.ErrNoDocuments
	}
	return room, nil
}

func GetMinimalizedRoom(roomId string) (Room, error) {
	var room Room
	// use $lookup to get the User object for RoomOwner
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
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
		{Key: "password", Value: 0},
		{Key: "room_owner.password", Value: 0},
		{Key: "room_owner.rooms", Value: 0},
		{Key: "messages", Value: 0},
		{Key: "users", Value: 0},
	}}}
	// use $sort to sort messages by createdAt in descending order

	// use $limit to limit the number of messages to 50
	// aggregate the pipeline stages

	pipeline := mongo.Pipeline{matchStage, lookupStage, unwindStage, projectStage}
	cursor, err := GetRoomCollection().Aggregate(context.Background(), pipeline)
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
	return room, nil
}

func IsUserJoinedInTheRoom(username string, roomId string) (room Room, isJoined bool) {
	var r Room
	r, err := GetRoom(roomId)
	fmt.Println(r.Users)
	room = r
	fmt.Println(room, err)
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
