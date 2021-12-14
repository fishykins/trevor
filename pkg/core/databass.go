package core

import (
	"context"

	"github.com/fishykins/trevor/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Databass struct {
	client *mongo.Client
}

func (d *Databass) Connect() error {
	err := d.client.Connect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (d *Databass) Disconnect() error {
	err := d.client.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (d *Databass) Internal() *mongo.Database {
	return d.client.Database("trevor")
}

func (d *Databass) Dictionary() *mongo.Database {
	return d.client.Database("babelfish")
}

func (d *Databass) EnglishDictionary() *mongo.Collection {
	return d.Dictionary().Collection("english")
}

func (d *Databass) Users() []*models.User {
	collection := d.Internal().Collection("users")
	ctx := context.Background()
	res, err := collection.Find(ctx, bson.D{})
	if err != nil {
		Error("Error getting users from mongo:", err)
		return nil
	}
	var users []*models.User
	if err := res.All(ctx, &users); err != nil {
		Error("Error parsing users from mongo:", err)
	}
	for _, user := range users {
		user.Init()
	}
	return users
}

func (d *Databass) User(id uint64) *models.User {
	collection := d.Internal().Collection("users")
	ctx := context.Background()
	filter := bson.M{"_id": id}
	res := collection.FindOne(ctx, filter)

	var user models.User
	if err := res.Decode(&user); err != nil {
		Error("Error decoding user from mongo:", err)
		return nil
	}
	user.Init()
	return &user
}

func (d *Databass) UpdateUser(ctx *context.Context, user *models.User) error {
	collection := d.Internal().Collection("users")
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

func (d *Databass) RemoveUser(ctx *context.Context, user *models.User) error {
	collection := d.Internal().Collection("users")
	filter := bson.M{"_id": user.DiscordID}
	collection.FindOneAndDelete(*ctx, filter)
	return nil
}
