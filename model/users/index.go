package users

import (
	"context"
	"time"

	"github.com/degreane/ezekod.com/middleware/ezelogger"
	"github.com/degreane/ezekod.com/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Group     string             `json:"group" bson:"group"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

var (
	Users mongo.Collection
)

func (u *User) Insert() {
	res, err := model.DB.DataBase.Collection("users").InsertOne(context.TODO(), u)
	if err != nil {
		ezelogger.Ezelogger.Fatalf("%+v", err)
	}
	ezelogger.Ezelogger.Printf("%+v", res)

}

func TestInsertUser() {
	user := User{
		UserName:  "fbanna",
		Password:  "shta2telik",
		FirstName: "Faysal",
		LastName:  "Banna",
		Group:     "Admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user.Insert()
}
