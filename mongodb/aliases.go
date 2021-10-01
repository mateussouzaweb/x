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

// Document struct
type Document = interface{}

// Results struct
type Results = []interface{}

// == REPRESENTATIONS ==

// List struct
type List = bson.A

// Array struct
type Array = bson.D

// Item struct
type Item = bson.E

// Map struct
type Map = bson.M

// Pipeline struct
type Pipeline = []Array

// == PRIMITIVES ==

// Binary struct
type Binary = primitive.Binary

// DateTime struct
type DateTime = primitive.DateTime

// Decimal128 struct
type Decimal128 = primitive.Decimal128

// JavaScript struct
type JavaScript = primitive.JavaScript

// MaxKey struct
type MaxKey = primitive.MaxKey

// MinKey struct
type MinKey = primitive.MinKey

// Null struct
type Null = primitive.Null

// ObjectID struct
type ObjectID = primitive.ObjectID

// Regex struct
type Regex = primitive.Regex

// Timestamp struct
type Timestamp = primitive.Timestamp

// == OPTIONS ==

// Collation struct
type Collation = options.Collation

// CursorType struct
type CursorType = options.CursorType
