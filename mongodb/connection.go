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
	URI                   string
	Database              string
	OperationTimeout      int64
	MassOperationTimeout  int64
	IndexOperationTimeout int64
}

// Context method
func Context(timeout int64) (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Second,
	)
}

// Connect method
func Connect(config *Config) error {

	ctx, cancel := Context(config.OperationTimeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(ctx, clientOptions)

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

// Disconnect method
func Disconnect() error {
	ctx := context.Background()
	return _client.Disconnect(ctx)
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
