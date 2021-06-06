package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	pointsCollection *mongo.Collection
)

type Points struct {
	ChatID int64 `bson:"chat_id"`
	UserID int64 `bson:"user_id"`
	Points int64 `bson:"points"`
}

func getPointsCollection() *mongo.Collection {
	if pointsCollection != nil {
		return pointsCollection
	}
	pointsCollection = GetDatabase().Collection("points")
	return pointsCollection
}

func UpdatePoints(chatID int64, userID int64, actionSign string) error {
	points := 1
	if actionSign == "-" {
		points = -1
	}

	_, err := getPointsCollection().UpdateOne(
		Ctx,
		bson.M{
			"chat_id": chatID,
			"user_id": userID,
		},
		bson.M{
			"$inc": bson.M{"points": points},
		},
		options.Update().SetUpsert(true),
	)
	return err
}

func GetPoints(chatID int64, userID int64) (Points, error) {
	var points Points
	err := getPointsCollection().FindOne(
		Ctx,
		bson.M{
			"chat_id": chatID,
			"user_id": userID,
		},
	).Decode(&points)
	return points, err
}

func GetTopPoints(chatID int64, isPositive bool) ([]Points, error) {
	var result []Points
	opts := options.Find()
	if isPositive {
		opts.SetSort(bson.M{"points": 1})
	} else {
		opts.SetSort(bson.M{"points": -1})
	}

	curr, err := getPointsCollection().Find(
		Ctx,
		bson.M{"chat_id": chatID},
		opts,
	)
	if err != nil {
		return result, err
	}

	err = curr.All(Ctx, &result)
	return result, err
}

func GetUserPoints(chatID int64, userID int64) (Points, error) {
	var result Points
	err := getPointsCollection().FindOne(
		Ctx,
		bson.M{
			"chat_id": chatID,
			"user_id": userID,
		},
	).Decode(&result)
	return result, err
}
