package main

import (
	"context"
	"errors"
	"fmt"
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
	Delete(ctx context.Context, id string)error
	GetApplicationById(ctx context.Context,id string)(*PostApplication,error)
	UpdateApplication(ctx context.Context, application map[string]interface{},id string)error
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
func (store *mongoStore)Delete(ctx context.Context ,id string)error{
	filter:=bson.M{"id":id}
	coll:=store.client.Database("exam-application").Collection("application")
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
	coll:=store.client.Database("exam-application").Collection("application")
	filters:=bson.M{"id":id}
	result:=coll.FindOne(ctx,filters)
	err:=result.Decode(&application)
	if err!=nil{
		return nil,err
	}
	return &application,nil
}
func(store *mongoStore)UpdateApplication(ctx context.Context,application map[string]interface{},id string)error{
	coll:=store.client.Database("exam-application").Collection("application")
	filters:=bson.M{"id":id}
	fields:= bson.M{}
	for k,v:=range application{
		fields[k] = v
	}

	update:=bson.M{
		"$set":fields,
	}
	_,err:=coll.UpdateOne(ctx,filters,update)
	// if result.==0{
	// 	return errors.New("the id of the  doucment to be updated does not exists")
	// }
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
		
	return &mongoStore{
		client: client,
	},nil
}


