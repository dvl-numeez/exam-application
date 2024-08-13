package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)




type Storage interface {

}

type mongoStore struct{
	client *mongo.Client
}

func NewMongoStore(ctx context.Context) (*mongoStore,error){
	godotenv.Load()
	mongoUrl:=os.Getenv("MONGO_URL")
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI(mongoUrl))
	if err !=nil{
		return nil,err
	}
	if err:=client.Ping(ctx,readpref.Primary());err!=nil{
		return nil,err
	}
		
	return &mongoStore{
		client: client,
	},nil
}