package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(db, collection string) *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error connect to Mongodb", err)
	}
	defer cancel()
	coll := client.Database(db).Collection(collection)
	return coll
}

func InsertOne(db, collection string, data interface{}) (*mongo.InsertOneResult, error) {
	coll := Connect(db, collection)
	return coll.InsertOne(context.Background(), data)
}

func Find(db, collection string, query interface{}) ([]interface{}, error) {
	coll := Connect(db, collection)
	var v []interface{}
	cur, err := coll.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	err = cur.All(context.Background(), &v)
	return v, err
}

func FindOne(db, collection string, query interface{}) (interface{}, error) {
	coll := Connect(db, collection)
	var v interface{}
	err := coll.FindOne(context.Background(), query).Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, err
}

/*
func FindAll(db, collection string) {
	coll := Connect(db, collection)
	return coll.FindAll(context.Background(), bson.D{})
}

func Remove(db, collection string, query interface{}) error {
	coll := Connect(db, collection)
	return coll.Remove(context.Background(), query)
}
*/
