package mongo_persistence

import (
	"context"
	"fmt"
	"musical-tables-api/services/room-admin/internal/domain"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbRepository struct {
	db *mongo.Database
}

// NewMongoClient creates a new MongoClient instance
func NewMongoDbRepository() (*MongoDbRepository, error) {
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

	db := client.Database("rooms")

	return &MongoDbRepository{db: db}, nil
}

func (c *MongoDbRepository) Close(ctx context.Context) error {
	return c.db.Client().Disconnect(ctx)
}

func (c *MongoDbRepository) SaveRoom(ctx context.Context, room *domain.Room) error {
	collection := c.db.Collection("rooms")

	filter := bson.M{"_id": room.ID()}

	roomDoc := bson.M{
		"_id":    room.ID(),
		"name":   room.Name(),
		"tables": room.Tables(),
	}
	//_, err := collection.InsertOne(ctx, roomDoc)
	_, err := collection.ReplaceOne(ctx, filter, roomDoc, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error creating room: %v", err)
	}

	return nil
}

func (c *MongoDbRepository) GetRoomByID(ctx context.Context, id string) (*domain.Room, error) {
	// Define filter to find the room with the given ID
	filter := bson.M{"_id": id}

	// Find the room in the database
	var room domain.Room
	err := c.db.Collection("rooms").FindOne(ctx, filter).Decode(&room)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrRoomNotFound
		}
		return nil, err
	}

	return &room, nil
}
