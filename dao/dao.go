package dao

import (
	"apis/config"
	dto "apis/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDao struct {
	collection *mongo.Collection
}

func NewMongoDao() MongoDao {
	clientOptions := options.Client().ApplyURI(config.MONGO_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	return MongoDao{
		collection: client.Database("vivek_db").Collection("hell_and_heven"),
	}
}

func (m *MongoDao) GetAllUsers() ([]dto.MainDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	var users []dto.MainDB
	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}

	return users, nil
}

func (m *MongoDao) GetUserByName(name string) (dto.MainDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user dto.MainDB
	filter := bson.M{"name": name}
	err := m.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *MongoDao) CreateUser(mainDB dto.MainDB) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.collection.InsertOne(ctx, mainDB)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MongoDao) UpdateUser(name string, updatedData dto.MainDB) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"name": name}
	update := bson.M{
		"$set": updatedData,
	}

	result, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MongoDao) DeleteUser(name string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"name": name}
	result, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
