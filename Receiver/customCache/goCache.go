package customCache

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Db *mongo.Database

func (m *Message) Update(ctx context.Context, db *mongo.Database, collectionName string, filter Message, update Message) error {
	collection := db.Collection(collectionName)

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) Read(ctx context.Context, db *mongo.Database, collectionName string, filter Message, result Message) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) ReadAll(ctx context.Context, db *mongo.Database, collectionName string, result []Message) error {
	collection := db.Collection(collectionName)

	cursor, err := collection.Find(ctx, nil)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, result); err != nil {
		return err
	}

	return nil
}

func (m *Message) Create(ctx context.Context, db *mongo.Database, collectionName string, model Message) error {
	collection := db.Collection(collectionName)

	_, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}
