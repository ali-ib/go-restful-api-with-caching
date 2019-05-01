package dao

import (
	"context"
	"go-restful-api-with-caching/configs"
	"go-restful-api-with-caching/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var ctx context.Context
var config configs.Config

// Initializing and establishing a connection to database
func init() {
	config = configs.Configs
	ctx = context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(config.CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(config.DBNAME)
}

// Returns resources filtered by 'filter'
func Find(filter bson.M) ([]models.Person, error) {
	cur, err := db.Collection(config.COLLECTIONNAME).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var elements []models.Person
	var elem models.Person
	// Get the next result from the cursor
	for cur.Next(ctx) {
		if err := cur.Decode(&elem); err != nil {
			return nil, err
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return elements, nil
}

// Insert new resource into database
func InsertOne(person models.Person) (primitive.ObjectID, error) {
	res, err := db.Collection(config.COLLECTIONNAME).InsertOne(ctx, bson.M{"name": person.Name, "age": person.Age})
	if err != nil {
		return primitive.NilObjectID, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

// Delete resource from database
func DeleteOne(personID string) (int64, error) {
	id, err := primitive.ObjectIDFromHex(personID)
	if err != nil {
		return 0, err
	}
	res, err := db.Collection(config.COLLECTIONNAME).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

// Update resource in database
func UpdateOne(person models.Person, personID string) (int64, error) {
	id, err := primitive.ObjectIDFromHex(personID)
	if err != nil {
		return 0, err
	}
	res, err := db.Collection(config.COLLECTIONNAME).UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"name": person.Name, "age": person.Age}},
	)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}
