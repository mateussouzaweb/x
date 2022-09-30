package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ListIndexesData struct
type ListIndexesData struct {
	Collection string
}

// ListIndexes method
func ListIndexes(data ListIndexesData, destination interface{}) error {

	ctx, cancel := Context(_config.IndexOperationTimeout)
	defer cancel()

	indexes := GetCollection(data.Collection).Indexes()
	cursor, err := indexes.List(ctx, &options.ListIndexesOptions{})

	if err != nil {
		return err
	}

	if err := cursor.All(ctx, destination); err != nil {
		return err
	}

	return nil
}

// IndexOptions struct
type IndexOptions = options.IndexOptions

// IndexData struct
type IndexData struct {
	Collection string
	Name       string
	Keys       Document
	Options    *IndexOptions
}

// EnsureIndexes method
func EnsureIndexes(data []IndexData) error {

	ctx, cancel := Context(_config.IndexOperationTimeout)
	defer cancel()

	for _, item := range data {
		item.Options.Name = &item.Name
		indexes := GetCollection(item.Collection).Indexes()

		_, err := indexes.CreateOne(ctx, mongo.IndexModel{
			Keys:    item.Keys,
			Options: item.Options,
		}, &options.CreateIndexesOptions{})

		if err != nil {
			return err
		}
	}

	return nil
}

// DropIndexes method
func DropIndexes(data []IndexData) error {

	ctx, cancel := Context(_config.IndexOperationTimeout)
	defer cancel()

	for _, item := range data {
		indexes := GetCollection(item.Collection).Indexes()
		_, err := indexes.DropOne(ctx, item.Name, &options.DropIndexesOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
