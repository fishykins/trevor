package core

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
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

func (d *Databass) Users() *mongo.Collection {
	return d.Internal().Collection("users")
}
