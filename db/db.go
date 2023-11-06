package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://emi:naog7412@cluster0.ouboxvu.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		panic(err)
	}

	return client
}