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
	t.Run("Deleting an actual document",func(t *testing.T){
		err:=store.Delete(context.TODO(),"5749302")	
		if err!=nil{
			t.Error("not expecting an error but got an error")
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