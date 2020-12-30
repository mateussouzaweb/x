package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var _client *Client
var _database *Database
var _config *Config

// Config struct
type Config struct {
	URI      string
	Database string
}

// Context method
func Context(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

// Connect method
func Connect(config *Config) error {

	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		return err
	}

	ctx, cancel := Context(10 * time.Second)
	defer cancel()

	err = client.Connect(ctx)

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return err
	}

	_client = client
	_database = client.Database(config.Database)
	_config = config

	return nil
}

// GetClient method
func GetClient() *Client {
	return _client
}

// GetDatabase method
func GetDatabase() *Database {
	return _database
}

// GetCollection method
func GetCollection(name string) *Collection {
	return _database.Collection(name)
}
