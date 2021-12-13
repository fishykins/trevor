package models

import (
	"context"

	"github.com/fishykins/trevor/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        uint64    `json:"ID" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	Alignment Alignment `json:"alignment" bson:"alignment,omitempty"`
}

func (u User) PushToDatabass(ctx *context.Context) error {
	collection := core.App().Databass().Users()
	filter := bson.M{"_id": u.ID}

	update := bson.M{
		"$set": u,
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
