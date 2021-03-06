package dao

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/gutorc92/api-farm/collections"
)

type DataMongo struct {
	Uri string
	Client *mongo.Client
	Database string
}

func NewDataMongo (uri string, database string) (*DataMongo, error) {
	var dt DataMongo
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dt.Database = database
	dt.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Error to connect")
		return nil, err
	}
	err = dt.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error to connect")
		return nil, err
	}
	return &dt, nil
}

func (dt *DataMongo) InsertFarm(farm collections.Farm) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := dt.Client.Database(dt.Database).Collection("farm")
	res, err := collection.InsertOne(ctx, farm)
	if err != nil {
		fmt.Println("Error to error")
		return primitive.NewObjectID(), err
	}
	id := res.InsertedID.(primitive.ObjectID)
	// if err != nil {
	// 	return primitive.NewObjectID(), nil
	// }
	fmt.Println("Id insert: ", id)
	return id, nil
}

func (dt *DataMongo) Insert(collectionName string, save interface{}) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := dt.Client.Database(dt.Database).Collection(collectionName)
	res, err := collection.InsertOne(ctx, save)
	if err != nil {
		fmt.Println("Error to error")
		return primitive.NewObjectID(), err
	}
	id := res.InsertedID.(primitive.ObjectID)
	// if err != nil {
	// 	return primitive.NewObjectID(), nil
	// }
	fmt.Println("Id insert: ", id)
	return id, nil
}

func (dt *DataMongo) FindFarm() ([]collections.Farm, error) {
	var farms []collections.Farm
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := dt.Client.Database(dt.Database).Collection("farm")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &farms); err != nil {
		return nil, err
	}
	return farms, nil
}

func (dt *DataMongo) FindAll(collectionName string) ([]interface{}, error) {
	var data []interface{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := dt.Client.Database(dt.Database).Collection(collectionName)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

