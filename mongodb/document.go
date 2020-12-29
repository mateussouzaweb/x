package mongodb

import (
	"strings"
	"time"

	"github.com/google/uuid"
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

// Find method
func Find(results interface{}, document Document, filters *Array, options *Options) error {

	ctx, cancel := Context(10 * time.Second)
	defer cancel()

	collection := document.TheCollection()
	cursor, err := collection.Find(ctx, filters, options)

	if err != nil {
		return err
	}

	if err := cursor.All(ctx, results); err != nil {
		return err
	}

	return nil
}

// Retrieve method
func Retrieve(document Document, filters *Array) error {

	ctx, cancel := Context(10 * time.Second)
	defer cancel()

	collection := document.TheCollection()
	err := collection.FindOne(ctx, filters).Decode(document)

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

// Create method
func Create(document Document) (*InsertResult, error) {

	ctx, cancel := Context(5 * time.Second)
	defer cancel()

	collection := document.TheCollection()
	result, err := collection.InsertOne(ctx, document)

	return result, err
}

// Update method
func Update(document Document) (*UpdateResult, error) {

	ctx, cancel := Context(5 * time.Second)
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

// Delete method
func Delete(document Document) (*DeleteResult, error) {

	ctx, cancel := Context(5 * time.Second)
	defer cancel()

	collection := document.TheCollection()
	filters := Array{
		Item{Key: "uuid", Value: document.TheUUID()},
	}

	result, err := collection.DeleteOne(ctx, filters)

	return result, err
}
