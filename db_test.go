package main

import (
	"context"
	"errors"
	"testing"
	"time"
)


func TestInsertApplication(t *testing.T) {
		store,err:=NewMongoStore(context.TODO())
		if err!=nil{
			t.Errorf("cannot get the store")
		}
		application:=&Application{
			FirstName: "Numeez",
			LastName: "Baloch",
			MiddleName: "Khan",
			FatherFirstName: "Asif",
			FatherMiddleName: "Khan",
			FatherLastName: "Baloch",
			Gender: "male",
			BoardName: "Gujarat State Board",
			Pincode: 380051,
			City: "Ahmedabad",
			StateOfDomicile: "Gujarat",
			State: "Gujarat",
			Village: "Modasa",
			RollNumber: "355678692",
			YearOfPassing: "2023",
			DOB: time.Date(2001,time.September,24,0,0,0,0,time.Local),
			District: "Ahmedabad",
			HomeDistrict: "Ahmedabad",
			Address: "1 sterling heights",
			HouseNo: "17",
		}

		err=store.InsertApplication(context.TODO(),application)
		if err!=nil{
			t.Errorf("error in inserting application")
		}

		t.Run("Entering invalid gender",func(t *testing.T){
			application:=&Application{
				FirstName: "Numeez",
				LastName: "Baloch",
				MiddleName: "Khan",
				FatherFirstName: "Asif",
				FatherMiddleName: "Khan",
				FatherLastName: "Baloch",
				Gender: "abc",
				BoardName: "Gujarat State Board",
				Pincode: 380051,
				City: "Ahmedabad",
				StateOfDomicile: "Gujarat",
				State: "Gujarat",
				Village: "Modasa",
				RollNumber: "355678692",
				YearOfPassing: "2023",
				DOB: time.Date(2001,time.September,24,0,0,0,0,time.Local),
				District: "Ahmedabad",
				HomeDistrict: "Ahmedabad",
				Address: "1 sterling heights",
				HouseNo: "17",
			}
			expected:=errors.New("only male and female genders are acceptable")
			got:=store.InsertApplication(context.TODO(),application)
			if got.Error()!=expected.Error(){
				t.Errorf("Actual error : %s Expected error : %s",got,expected)
			}
	
			
		})
}

func TestDelete(t *testing.T){
	store,err:=NewMongoStore(context.TODO())
		if err!=nil{
			t.Errorf("cannot get the store")
		}
	t.Run("Deleting document with wrong ID",func(t *testing.T){
		err:=store.Delete(context.TODO(),"5749302")	
		expected:=errors.New("the id referencing to the document in not correct such document does not exists")
		if err.Error()!=expected.Error(){
			t.Errorf("Actual error : %s Expected error : %s ",err.Error(),expected.Error())
		}

	})
	// this test will run for once because it actually deletes the data from the database since we are not using mocks
	// that is why it is commented
	// t.Run("Deleting an actual document",func(t *testing.T){
	// 	err:=store.Delete(context.TODO(),"4a1fed12-47a4-4220-bdad-779b0fd6b80f")	
	// 	if err!=nil{
	// 		t.Error("not expecting an error but got an error")
	// 	}
	// })
}

func TestUpdateApplication(t *testing.T){
	store,err:=NewMongoStore(context.TODO())
	if err!=nil{
		t.Errorf("cannot get the store")
	}
	t.Run("Updating an application with the wrong id",func(t *testing.T) {
		err:=store.UpdateApplication(context.TODO(),Data{},"123456765")
		if err==nil{
			t.Error("Expected an error but did not get it ")
		}
	})
	t.Run("Updating an application", func(t *testing.T){
		id:="11946ff2-c7e5-4942-9e43-630107058da7"
		firstName:="Numeez Khan"
		middleName:="Asif"
		lastName:="Baloch"
		data:=Data{
			"firstname":firstName,
			"middlename":middleName,
			"lastname":lastName,
		}
		err:=store.UpdateApplication(context.TODO(),data,id)
		if err!=nil{
			t.Error("Error occured while upating the application : ",err)
		}
		
	})
}

func TestGetApplictionById(t *testing.T){
	store,err:=NewMongoStore(context.TODO())
		if err!=nil{
			t.Errorf("cannot get the store")
		}
	t.Run("Testing by inputing valid id",func(t *testing.T){
		wantedId:="11946ff2-c7e5-4942-9e43-630107058da7"
		application,err:=store.GetApplicationById(context.TODO(),wantedId)
		if err!=nil{
			t.Error("Error occured while fetching the application : ",err)
		}
		got:=application.Id
		if got!=wantedId{
			t.Errorf("Actual : %s Expected : %s",got,wantedId)
		}
		
	})
	t.Run("Testing by inputing invalid id",func(t *testing.T){
		id:="11946ff2-12345gh"
		_,err:=store.GetApplicationById(context.TODO(),id)
		if err==nil{
			t.Errorf("Wanted an error but did not give an error")
		}
		
	})
}



func BenchmarkTest(b *testing.B) {
	store,err:=NewMongoStore(context.TODO())
		if err!=nil{
			b.Errorf("cannot get the store")
		}
	for i:=0;i<b.N;i++{
		err:=store.Delete(context.TODO(),"5749302")	
		expected:=errors.New("the id referencing to the document in not correct such document does not exists")
		if err.Error()!=expected.Error(){
			b.Errorf("Actual error : %s Expected error : %s ",err.Error(),expected.Error())
		}
	}
}