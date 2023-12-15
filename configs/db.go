package configs

import (
	"context"
	"fmt"
	"graphql-server/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func ConnectDB() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return &DB{client: client}
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("ipl2023").Collection(collectionName)
}

func (db *DB) CreateData(input *model.NewData) (*model.Data, error) {
	collection := colHelper(db, "data")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	properties := model.JSON(input.Properties)
	data := &model.Data{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		ProjectID:   input.ProjectID,
		UserID:      input.UserID,
		AccountID:   input.AccountID,
		EventName:   input.EventName,
		ProjectName: input.ProjectName,
		Properties:  properties,
	}

	return data, err
}

func (db *DB) GetDataList() ([]*model.Data, error) {
	collection := colHelper(db, "data")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var dataList []*model.Data
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleData *model.Data
		if err = res.Decode(&singleData); err != nil {
			log.Fatal(err)
		}
		dataList = append(dataList, singleData)
	}

	return dataList, err
}
