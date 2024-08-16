package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Data map[string]interface{}
//TODO
//For checking the fields are valid or not
// fields:=map[string]bool{}

type Storage interface {
	InsertApplication(ctx context.Context,appilcation *Application)error
	FetchAll(ctx context.Context,filters Data)([]PostApplication,error)
	Delete(ctx context.Context, id string)error
	GetApplicationById(ctx context.Context,id string)(*PostApplication,error)
	UpdateApplication(ctx context.Context, application Data,id string)error
}

type mongoStore struct{
	database *mongo.Database
}

func (store *mongoStore)InsertApplication(ctx context.Context,application *Application)error{
	if !application.ValidateGender() {
		return errors.New("only male and female genders are acceptable")
	}
	postApplication:=application.NewApplicationPost()
	postApplication.Gender = strings.ToLower(application.Gender)
	coll:=store.database.Collection("application")
	_,err:=coll.InsertOne(ctx,postApplication)
	if err!=nil{
		return err
	}

	return nil
}

func(store *mongoStore)FetchAll(ctx context.Context,filters Data)([]PostApplication,error){
	coll:=store.database.Collection("application")
	bsonFilters:=makeBson(filters)
	cursor,err:=coll.Find(ctx,bsonFilters)
	if err!=nil{
		return nil,err
	}
	var applications []PostApplication
	if err:=cursor.All(ctx,&applications);err!=nil{
		return nil,err
	}
	if applications==nil{
		return nil,errors.New("the filters you provide does not exists check your fields again")
	}
	return applications,nil
}
func (store *mongoStore)Delete(ctx context.Context ,id string)error{
	filter:=bson.M{"id":id}
	coll:=store.database.Collection("application")
	deletedDocumentNum,err:=coll.DeleteOne(ctx,filter)
	if deletedDocumentNum.DeletedCount==0{
		return errors.New("the id referencing to the document in not correct such document does not exists")
	}
	if err!=nil{
		return err
	}
	return nil
}

func(store *mongoStore)GetApplicationById(ctx context.Context,id string)(*PostApplication,error){
	var application PostApplication
	coll:=store.database.Collection("application")
	filters:=bson.M{"id":id}
	result:=coll.FindOne(ctx,filters)
	err:=result.Decode(&application)
	if err!=nil{
		return nil,err
	}
	return &application,nil
}
func(store *mongoStore)UpdateApplication(ctx context.Context,filters Data,id string)error{
	coll:=store.database.Collection("application")
	_,err:=store.GetApplicationById(ctx,id)
	if err!=nil{
		return err
	}
	options:=bson.M{"id":id}
	fields:=makeBson(filters)
	update:=bson.M{
		"$set":fields,
	}
	_,err=coll.UpdateOne(ctx,options,update)
	if err!=nil{
		return err
	}
	
	
	
	return nil
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
	fmt.Println("Connected to the database")
	db:=client.Database("exam-application")	
	return &mongoStore{
		database: db,
	},nil
}


func makeBson(filters Data)bson.M{
	result:=bson.M{}
	for k,v:=range filters{
		lowerCaseKey := strings.ToLower(k)
		result[lowerCaseKey] = v
	}
	return result 
}
