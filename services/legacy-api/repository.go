package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient represents the MongoDB client
type MongoClient struct {
	db *mongo.Database
}

// NewMongoClient creates a new MongoClient instance
func NewMongoClient() (*MongoClient, error) {
	mongoHost := os.Getenv("MONGO_HOST")
	if mongoHost == "" {
		mongoHost = "localhost"
	}
	mongoPort := os.Getenv("MONGO_PORT")
	if mongoPort == "" {
		mongoPort = "27017"
	}
	mongoURI := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("restaurant")

	return &MongoClient{db: db}, nil
}

func (c *MongoClient) Close(ctx context.Context) error {
	return c.db.Client().Disconnect(ctx)
}

func (c *MongoClient) SaveRoom(ctx context.Context, room *Room) error {
	collection := c.db.Collection("rooms")

	filter := bson.M{"_id": room.ID}
	_, err := collection.ReplaceOne(ctx, filter, room, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error creating room: %v", err)
	}

	return nil
}

func (c *MongoClient) SaveTable(ctx context.Context, table *Table) error {
	collection := c.db.Collection("tables")

	filter := bson.M{"_id": table.ID}
	_, err := collection.ReplaceOne(ctx, filter, table, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error creating tabel: %v", err)
	}

	return nil
}

func (c *MongoClient) GetRoomByID(ctx context.Context, id string) (*Room, error) {
	// Define filter to find the room with the given ID
	filter := bson.M{"_id": id}

	// Find the room in the database
	var room Room
	err := c.db.Collection("rooms").FindOne(ctx, filter).Decode(&room)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrRoomNotFound
		}
		return nil, err
	}

	return &room, nil
}

func (c *MongoClient) GetTableByID(ctx context.Context, id string) (*Table, error) {
	filter := bson.M{"_id": id}

	var table Table
	err := c.db.Collection("tables").FindOne(ctx, filter).Decode(&table)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrTableNotFound
		}
		return nil, err
	}

	return &table, nil
}

var ErrRoomNotFound = errors.New("room not found")
var ErrTableNotFound = errors.New("table not found")
