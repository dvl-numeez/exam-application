package main

import (
	"context"
	"errors"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)




type Storage interface {
	InsertApplication(ctx context.Context,appilcation *Application)error
	FetchAll(ctx context.Context)([]PostApplication,error)
}

type mongoStore struct{
	client *mongo.Client
}

func (store *mongoStore)InsertApplication(ctx context.Context,application *Application)error{
	if !application.ValidateGender() {
		return errors.New("only male and female genders are acceptable")
	}
	postApplication:=application.NewApplicationPost()
	coll:=store.client.Database("exam-application").Collection("application")
	_,err:=coll.InsertOne(ctx,postApplication)
	if err!=nil{
		return err
	}

	return nil
}

func(store *mongoStore)FetchAll(ctx context.Context)([]PostApplication,error){
	coll:=store.client.Database("exam-application").Collection("application")
	cursor,err:=coll.Find(ctx,bson.D{})
	if err!=nil{
		return nil,err
	}
	var applications []PostApplication
	if err:=cursor.All(ctx,&applications);err!=nil{
		return nil,err
	}
	
	return applications,nil
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


