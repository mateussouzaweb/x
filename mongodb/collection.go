package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOptions struct
type FindOptions = options.FindOptions

// FindData struct
type FindData struct {
	Collection string
	Filters    *Array
	Options    *FindOptions
}

// Find method
func Find(data FindData, destination any) error {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	cursor, err := collection.Find(ctx, data.Filters, data.Options)

	if err != nil {
		return err
	}

	if err := cursor.All(ctx, destination); err != nil {
		return err
	}

	return nil
}

// FindOneOptions struct
type FindOneOptions = options.FindOneOptions

// FindOneData struct
type FindOneData struct {
	Collection string
	Filters    *Array
	Options    *FindOneOptions
}

// FindOne method
func FindOne(data FindOneData, destination Document) error {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	err := collection.FindOne(ctx, data.Filters, data.Options).Decode(destination)

	if err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}

// FindOneByData struct
type FindOneByData struct {
	Collection string
	Key        string
	Value      string
	Options    *FindOneOptions
}

// FindOneBy method
func FindOneBy(data FindOneByData, destination Document) error {

	filters := Array{
		Item{Key: data.Key, Value: data.Value},
	}

	return FindOne(FindOneData{
		Collection: data.Collection,
		Filters:    &filters,
		Options:    data.Options,
	}, destination)
}

// CountOptions struct
type CountOptions = options.CountOptions

// CountData struct
type CountData struct {
	Collection string
	Filters    *Array
	Options    *CountOptions
}

// Count method
func Count(data CountData) (int64, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	count, err := collection.CountDocuments(ctx, data.Filters, data.Options)

	return count, err
}

// AggregateOptions struct
type AggregateOptions = options.AggregateOptions

// AggregateData struct
type AggregateData struct {
	Collection string
	Pipeline   Pipeline
	Options    *AggregateOptions
}

// Aggregate method
func Aggregate(data AggregateData, destination any) error {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	cursor, err := collection.Aggregate(ctx, data.Pipeline, data.Options)

	if err != nil {
		return err
	}

	if err := cursor.All(ctx, destination); err != nil {
		return err
	}

	return nil
}

// DistinctOptions struct
type DistinctOptions = options.DistinctOptions

// DistinctData struct
type DistinctData struct {
	Collection string
	Field      string
	Filters    *Array
	Options    *DistinctOptions
}

// Distinct method
func Distinct(data DistinctData) ([]any, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	results, err := collection.Distinct(ctx, data.Field, data.Filters, data.Options)

	return results, err
}

// InsertManyResult struct
type InsertManyResult = mongo.InsertManyResult

// InsertManyOptions struct
type InsertManyOptions = options.InsertManyOptions

// InsertManyData struct
type InsertManyData struct {
	Collection string
	Documents  []Document
	Options    *InsertManyOptions
}

// InsertMany method
func InsertMany(data InsertManyData) (*InsertManyResult, error) {

	ctx, cancel := Context(_config.MassOperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	result, err := collection.InsertMany(ctx, data.Documents, data.Options)

	return result, err
}

// InsertResult struct
type InsertResult = mongo.InsertOneResult

// InsertOptions struct
type InsertOptions = options.InsertOneOptions

// InsertData struct
type InsertData struct {
	Collection string
	Document   Document
	Options    *InsertOptions
}

// Insert method
func Insert(data InsertData) (*InsertResult, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	result, err := collection.InsertOne(ctx, data.Document, data.Options)

	return result, err
}

// UpdateManyResult struct
type UpdateManyResult = mongo.UpdateResult

// UpdateManyOptions struct
type UpdateManyOptions = options.UpdateOptions

// UpdateManyData struct
type UpdateManyData struct {
	Collection string
	Filters    *Array
	Updates    Document
	Options    *UpdateManyOptions
}

// UpdateMany method
func UpdateMany(data UpdateManyData) (*UpdateManyResult, error) {

	ctx, cancel := Context(_config.MassOperationTimeout)
	defer cancel()

	updates := Map{
		"$set": data.Updates,
	}

	collection := GetCollection(data.Collection)
	result, err := collection.UpdateMany(ctx, data.Filters, updates, data.Options)

	return result, err
}

// UpdateOneResult struct
type UpdateResult = mongo.UpdateResult

// UpdateOptions struct
type UpdateOptions = options.UpdateOptions

// UpdateData struct
type UpdateData struct {
	Collection string
	Filters    *Array
	Updates    Document
	Options    *UpdateOptions
}

// Update method
func Update(data UpdateData) (*UpdateResult, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	updates := Map{
		"$set": data.Updates,
	}

	collection := GetCollection(data.Collection)
	result, err := collection.UpdateOne(ctx, data.Filters, updates, data.Options)

	return result, err
}

// DeleteManyResult struct
type DeleteManyResult = mongo.DeleteResult

// DeleteManyOptions struct
type DeleteManyOptions = options.DeleteOptions

// DeleteManyData struct
type DeleteManyData struct {
	Collection string
	Filters    *Array
	Options    *DeleteManyOptions
}

// DeleteMany method
func DeleteMany(data DeleteManyData) (*DeleteManyResult, error) {

	ctx, cancel := Context(_config.MassOperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	result, err := collection.DeleteMany(ctx, data.Filters, data.Options)

	return result, err
}

// DeleteResult struct
type DeleteResult = mongo.DeleteResult

// DeleteOptions struct
type DeleteOptions = options.DeleteOptions

// DeleteData struct
type DeleteData struct {
	Collection string
	Filters    *Array
	Options    *DeleteOptions
}

// Delete method
func Delete(data DeleteData) (*DeleteResult, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := GetCollection(data.Collection)
	result, err := collection.DeleteOne(ctx, data.Filters, data.Options)

	return result, err
}
