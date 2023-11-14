package db_test

import (
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/test"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	teardown := test.SetupEssential()
	defer teardown()
	client := db.Connect(os.Getenv("MONGODB_URI"), "test")
	if client == nil {
		t.Error("Expected client to not be nil")
	}

}

func TestGetDB(t *testing.T) {
	destroyDb := test.SetupDBForTest()
	defer destroyDb()
	db := db.GetDB()
	if db == nil {
		t.Error("Expected db to not be nil")
	}
}

func TestDropDB(t *testing.T) {
	destroyDb := test.SetupDBForTest()
	defer destroyDb()
	err := db.DropDB()
	if err != nil {
		t.Error("Expected err to be nil but got ", err)
	}

}
