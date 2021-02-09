package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOptions struct
type FindOptions = options.FindOptions

// FindData struct
type FindData struct {
	Scope   Document
	Filters *Array
	Options *FindOptions
}

// Find method
func Find(results interface{}, data FindData) error {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := data.Scope.TheCollection()
	cursor, err := collection.Find(ctx, data.Filters, data.Options)

	if err != nil {
		return err
	}

	if err := cursor.All(ctx, results); err != nil {
		return err
	}

	return nil
}

// CountOptions struct
type CountOptions = options.CountOptions

// CountData struct
type CountData struct {
	Scope   Document
	Filters *Array
	Options *CountOptions
}

// Count method
func Count(data CountData) (int64, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := data.Scope.TheCollection()
	count, err := collection.CountDocuments(ctx, data.Filters, data.Options)

	return count, err
}

// AggregateOptions struct
type AggregateOptions = options.AggregateOptions

// AggregateData struct
type AggregateData struct {
	Scope    Document
	Pipeline *Pipeline
	Options  *AggregateOptions
}

// Aggregate method
func Aggregate(results interface{}, data AggregateData) error {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := data.Scope.TheCollection()
	cursor, err := collection.Aggregate(ctx, *data.Pipeline, data.Options)

	if err != nil {
		return err
	}

	if err := cursor.All(ctx, results); err != nil {
		return err
	}

	return nil
}

// DistinctOptions struct
type DistinctOptions = options.DistinctOptions

// DistinctData struct
type DistinctData struct {
	Scope   Document
	Field   string
	Filters *Array
	Options *DistinctOptions
}

// Distinct method
func Distinct(data DistinctData) ([]interface{}, error) {

	ctx, cancel := Context(_config.OperationTimeout)
	defer cancel()

	collection := data.Scope.TheCollection()
	results, err := collection.Distinct(ctx, data.Field, data.Filters, data.Options)

	return results, err
}

// CreateManyResult struct
type CreateManyResult = mongo.InsertManyResult

// CreateManyOptions struct
type CreateManyOptions = options.InsertManyOptions

// CreateManyData struct
type CreateManyData struct {
	Scope     Document
	Documents []interface{}
	Options   *CreateManyOptions
}

// CreateMany method
func CreateMany(data CreateManyData) (*CreateManyResult, error) {

	ctx, cancel := Context(_config.MassOperationTimeout)
	defer cancel()

	collection := data.Scope.TheCollection()
	result, err := collection.InsertMany(ctx, data.Documents, data.Options)

	return result, err
}

// UpdateManyResult struct
type UpdateManyResult = mongo.UpdateResult

// UpdateManyOptions struct
type UpdateManyOptions = options.UpdateOptions

// UpdateManyData struct
type UpdateManyData struct {
	Scope   Document
	Filters Array
	Update  Array
	Options *UpdateManyOptions
}

// UpdateMany method
func UpdateMany(data UpdateManyData) (*UpdateManyResult, error) {

	ctx, cancel := Context(_config.MassOperationTimeout)
	defer cancel()

	collection := data.Scope.TheCollection()
	result, err := collection.UpdateMany(ctx, data.Filters, data.Update)

	return result, err
}

// DeleteManyResult struct
type DeleteManyResult = mongo.DeleteResult

// DeleteManyOptions struct
type DeleteManyOptions = options.DeleteOptions

// DeleteManyData struct
type DeleteManyData struct {
	Scope   Document
	Filters *Array
	Options *DeleteManyOptions
}

// DeleteMany method
func DeleteMany(data DeleteManyData) (*DeleteManyResult, error) {

	ctx, cancel := Context(_config.MassOperationTimeout)
	defer cancel()

	collection := data.Scope.TheCollection()
	result, err := collection.DeleteMany(ctx, data.Filters, data.Options)

	return result, err
}
