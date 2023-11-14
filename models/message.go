package models

type Message struct {
	Id        string `json:"id" bson:"_id"`
	Username  string `json:"username" bson:"username"`
	Message   string `json:"message" bson:"message"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}
