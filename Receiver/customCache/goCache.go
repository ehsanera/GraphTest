package customCache

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Db *mongo.Database

func (m *Message) Update(ctx context.Context, db *mongo.Database, collectionName string, id primitive.ObjectID) error {
	collection := db.Collection(collectionName)

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", bson.D{{"received", true}}}})
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

func (m *Message) ReadAll(ctx context.Context, db *mongo.Database, collectionName string) ([]Message, error) {
	var result []Message

	collection := db.Collection(collectionName)

	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(context.TODO()) {
		var elem Message
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		result = append(result, elem)
	}

	return result, nil
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
