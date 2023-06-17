package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define Variables for the database
var (
	MONGODB_URI = "mongodb://psycho:shta2telik@127.0.0.1:27017/?authSource=admin"

	DB DBase
)

type DBase struct {
	Connection *mongo.Client
	DataBase   *mongo.Database
}

// DB returns a Database Client Connection Handler
func (c *DBase) Connect() error {
	// var err error
	db, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		return err
	} else {
		c.Connection = db
		c.DataBase = c.Connection.Database("ezekod")
		return nil
	}
}

func (c DBase) CloseDB() error {
	return c.Connection.Disconnect(context.Background())
}
func (c *DBase) SetDatabase(db string) {
	database := c.Connection.Database(db)
	c.DataBase = database
}
