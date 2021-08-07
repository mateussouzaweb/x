package mongodb

import (
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

// Document interface
type Document interface {
	TheUUID() string
	TheCollection() *Collection
}

// UUID method
func UUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// Retrieve method
func Retrieve(document Document, filters *Array) error {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := document.TheCollection()
	err := collection.FindOne(ctx, filters).Decode(document)

	if err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}

// RetrieveBy method
func RetrieveBy(document Document, key string, value *string) error {

	filters := Array{
		Item{Key: key, Value: value},
	}

	return Retrieve(document, &filters)
}

// RetrieveByUUID method
func RetrieveByUUID(document Document, uuid *string) error {
	return RetrieveBy(document, "uuid", uuid)
}

// CreateResult struct
type CreateResult = mongo.InsertOneResult

// Create method
func Create(document Document) (*CreateResult, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := document.TheCollection()
	result, err := collection.InsertOne(ctx, document)

	return result, err
}

// UpdateResult struct
type UpdateResult = mongo.UpdateResult

// Update method
func Update(document Document) (*UpdateResult, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := document.TheCollection()
	filters := Array{
		Item{Key: "uuid", Value: document.TheUUID()},
	}

	updates := Map{
		"$set": document,
	}

	result, err := collection.UpdateOne(ctx, filters, updates)

	return result, err
}

// DeleteResult struct
type DeleteResult = mongo.DeleteResult

// Delete method
func Delete(document Document) (*DeleteResult, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := document.TheCollection()
	filters := Array{
		Item{Key: "uuid", Value: document.TheUUID()},
	}

	result, err := collection.DeleteOne(ctx, filters)

	return result, err
}
