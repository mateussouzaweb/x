package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client struct
type Client = mongo.Client

// Database struct
type Database = mongo.Database

// Collection struct
type Collection = mongo.Collection

// Array struct
type Array = bson.D

// Item struct
type Item = bson.E

// Map struct
type Map = bson.M

// Pipeline struct
type Pipeline = []Array

// Regex struct
type Regex = primitive.Regex

// Options struct
type Options = options.FindOptions

// CountOptions struct
type CountOptions = options.CountOptions

// AggregateOptions struct
type AggregateOptions = options.AggregateOptions

// Collation struct
type Collation = options.Collation

// InsertResult struct
type InsertResult = mongo.InsertOneResult

// UpdateResult struct
type UpdateResult = mongo.UpdateResult

// DeleteResult struct
type DeleteResult = mongo.DeleteResult
