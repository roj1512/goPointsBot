package mongo

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usersCollection *mongo.Collection
)

type User struct {
	UserID    int64  `bson:"user_id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name,omitempty"`
}

func getUsersCollection() *mongo.Collection {
	if usersCollection != nil {
		return usersCollection
	}
	usersCollection = GetDatabase().Collection("users")
	return usersCollection
}

func UpdateUser(user gotgbot.User) error {
	_, err := getUsersCollection().UpdateOne(
		Ctx,
		bson.M{"user_id": user.Id},
		bson.M{
			"$set": User{
				user.Id,
				user.FirstName,
				user.LastName,
			},
		},
		options.Update().SetUpsert(true),
	)
	return err
}
