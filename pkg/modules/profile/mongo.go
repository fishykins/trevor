package profile

import (
	"context"

	"github.com/fishykins/trevor/pkg/core"
	"github.com/fishykins/trevor/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This is a set of functions to handle passing data back and forth to a mongo db. Should be handled directly by the Profile driver.

func GetUsersFromMongo() []*models.User {
	collection := core.App().Databass().Users()
	ctx := context.Background()
	res, err := collection.Find(ctx, bson.D{})
	if err != nil {
		core.Error("Error getting users from mongo:", err)
		return nil
	}
	var users []*models.User
	if err := res.All(ctx, &users); err != nil {
		core.Error("Error parsing users from mongo:", err)
	}
	return users
}

func UpdateMongoUser(ctx *context.Context, user *models.User) error {
	collection := core.App().Databass().Users()
	filter := bson.M{"_id": user.DiscordID}

	update := bson.M{
		"$set": user,
	}

	//Create an instance of an options and set the desired options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	//Find one result and update it
	collection.FindOneAndUpdate(*ctx, filter, update, &opt)
	return nil
}

func RemoveMongoUser(ctx *context.Context, user *models.User) error {
	collection := core.App().Databass().Users()
	filter := bson.M{"_id": user.DiscordID}
	collection.FindOneAndDelete(*ctx, filter)
	return nil
}
