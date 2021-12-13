package main

import (
	"context"

	"github.com/fishykins/trevor/pkg/core"
	"github.com/fishykins/trevor/pkg/modules/babel"
	"github.com/fishykins/trevor/pkg/modules/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//Background context
	ctx := context.Background()

	core.InitApplication("Mongo Test", []core.Module{babel.Babel, profile.Profile}, true)
	core.StartApplication()
	defer core.StopApplication()

	// testUser := models.User{
	// 	ID:        1234565,
	// 	Name:      "Fishykinz",
	// 	Alignment: models.ChaoticEvil(),
	// }

	users := core.App().Databass().Users()

	filter := bson.M{"_id": 1234565}

	update := bson.M{
		"$set": bson.M{"name": "Fishykinz"},
	}

	//Create an instance of an options and set the desired options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	//Find one result and update it
	users.FindOneAndUpdate(ctx, filter, update, &opt)

	//users.InsertOne(ctx, testUser)
}
