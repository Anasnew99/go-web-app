package models

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
	Description string   `json:"description" bson:"description"`
	Password    string   `json:"password" bson:"password"`
	RoomOwner   string   `json:"room_owner" bson:"room_owner"`
	CreatedAt   int64    `json:"created_at" bson:"created_at"`
	Id          string   `json:"id" bson:"_id"`
	Users       []string `json:"users" bson:"users"`
}
