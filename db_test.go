package main

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
		id:="2e78b0ea-82fc-42d8-82a0-65c08c891be8"
		firstName:="Numeez Khan"
		middleName:="Asif"
		lastName:="Baloch"
		data:=Data{
			"firstName":firstName,
			"middleName":middleName,
			"lastName":lastName,
		}
		err:=store.UpdateApplication(context.TODO(),data,id)
		if err!=nil{
			t.Error("Error occured while upating the application : ",err)
		}
		
	})
	t.Run("Updating fields that does not exists", func(t *testing.T){
		id:="2e78b0ea-82fc-42d8-82a0-65c08c891be8"
		firstName:="Numeez Khan"
		middleName:="Asif"
		lastName:="Baloch"
		data:=Data{
			"firstNme":firstName,
			"middleName":middleName,
			"lastNam":lastName,
		}
		err:=store.UpdateApplication(context.TODO(),data,id)
		got:=errors.New("the filters you provide does not exists check your fields again")
		if err.Error()!=got.Error(){
			t.Errorf("Actual error : %s Expected : %s",err.Error(),got.Error())
		}
		
	})
}

func TestGetApplictionById(t *testing.T){
	store,err:=NewMongoStore(context.TODO())
		if err!=nil{
			t.Errorf("cannot get the store")
		}
	t.Run("Testing by inputing valid id",func(t *testing.T){
		wantedId:="2e78b0ea-82fc-42d8-82a0-65c08c891be8"
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

func TestGetAllApplication(t *testing.T){
	store,err:=NewMongoStore(context.TODO())
		if err!=nil{
			t.Errorf("cannot get the store")
		}
	t.Run("Fetch All applications",func(t  *testing.T){
			applications,err:=store.FetchAll(context.TODO(),Data{})
			if err!=nil{
				t.Error("Error while fetching applications",err)
			}
			got:=len(applications)
			if got==0{
				t.Error("no applications in the database")
			}
	})
	t.Run("Fetch All applications with filter",func(t  *testing.T){
		applications,err:=store.FetchAll(context.TODO(),Data{"firstName":"Numeez"})
		if err!=nil{
			t.Error("Error while fetching applications",err)
		}
		length:=len(applications)
		if length==0{
			t.Error("Expected atleast one application with the filter")
		}
	})
	t.Run("Fetch All applications with wrong filters",func(t  *testing.T){
		_,err:=store.FetchAll(context.TODO(),Data{"college":"LJ"})
		if err==nil{
			t.Error("Expecting an error but did not get it")
		}
		
	})
}

func TestMakeBson(t *testing.T){
	data:=Data{
		"firstName":"Numeez",
		"lastName":"Baloch",
	}
	result:=makeBson(data)
	expected:=bson.M{
		"firstname":"Numeez",
		"lastname":"Baloch",
	}
	if !reflect.DeepEqual(result,expected){
		t.Errorf("Actual : %v, Expected : %v",result,expected)
	}
}
func TestCheckFields(t *testing.T){
	cases:= []struct{
		data Data
		result bool}{
		{data :Data{
			"firstName":"Numeez",
			"lastName":"Baloch",
		},
		result: true,	
	},
	{data :Data{
		"firstName":"Numeez",
		"last":"Baloch",
	},
	result: false,	
},
}
for _,c:= range cases{
	t.Run("Checking fields",func(t *testing.T) {
		actual:=checkFields(c.data)
		if actual!=c.result{
			t.Errorf("Actual : %t, Expected : %t",actual,c.result)
		}
	})
}
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