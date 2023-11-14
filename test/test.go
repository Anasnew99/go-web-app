package test

import (
	"anasnew99/server/chat_app/db"
	"os"

	"github.com/joho/godotenv"
)

func SetupEssential() func() error {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}

	return func() error {
		return nil
	}

}

func SetupDBForTest() func() error {
	tearEssential := SetupEssential()
	db.Connect(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_NAME"))

	return func() error {
		db.DropDB()
		db.Disconnect()
		tearEssential()
		return nil
	}
}
