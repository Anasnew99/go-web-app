package controllers

import (
	"anasnew99/server/chat_app/collections"
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func getMessageCollection() *mongo.Collection {
	return db.GetDB().Collection(collections.MESSAGES)
}

func addMessage(message models.Message) (*mongo.InsertOneResult, error) {

	// ignore id present in message

	newMessage, err := getMessageCollection().InsertOne(context.Background(), message)

	return newMessage, err
}
