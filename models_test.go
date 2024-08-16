package main

import (
	"testing"
	"time"
)

func TestValidateGender(t *testing.T) {
	applications := []Application{
		{Gender: "male"},
		{Gender: "Male"},
		{Gender: "MaLe"},
		{Gender: "female"},
		{Gender: "FEmale"},
		{Gender: "FEMALE"},
		{Gender: "MALE"},
		{Gender: "MALE"}}
	applicationsFail := []Application{
		{Gender: "abc"},
		{Gender: "queer"},
		{Gender: "def"},
		{Gender: "trans"},
		{Gender: "he"},
		{Gender: "she"},
		{Gender: "it"},
		{Gender: "idk"}}
	t.Run("Validate Gender function for right case", func(t *testing.T) {
		for _, application := range applications {
			got := application.ValidateGender()
			want := true
			if got != want {
				t.Errorf("Actual : %t Expected : %t", got, want)
			}
		}
	})
	t.Run("Validate Gender function for wrong case", func(t *testing.T) {
		for _, application := range applicationsFail {
			got := application.ValidateGender()
			want := false
			if got != want {
				t.Errorf("Actual : %t Expected : %t", got, want)
			}
		}
	})
}

func TestMakeFullName(t *testing.T) {
	type TestCase struct {
		application Application
		result      string
	}
	testCases := []TestCase{
		{Application{FirstName: "Numeez", MiddleName: "Khan", LastName: "Baloch"}, "Numeez Khan Baloch"},
		{Application{FirstName: "Asif", MiddleName: "Khan", LastName: "Baloch"}, "Asif Khan Baloch"},
		{Application{FirstName: "Abraham", MiddleName: "Benjamin", LastName: "Deviliers"}, "Abraham Benjamin Deviliers"},
	}

	t.Run("Testing the fullname", func(t *testing.T) {
		for _, test := range testCases {
			got := makeFullName(test.application.FirstName, test.application.MiddleName, test.application.LastName)
			want := test.result
			if got != want {
				t.Errorf("Actual %s Expected :%s", got, want)
			}
		}
	})
}

func TestCalculateAge(t *testing.T) {
	type TestCase struct{
		application Application
		result int
	}
	testCases:=[]TestCase{
		{Application{DOB:time.Date(2001,1,1,0,0,0,0,time.Local) },23},
		{Application{DOB:time.Date(2002,1,1,0,0,0,0,time.Local) },22},
		{Application{DOB:time.Date(2003,1,1,0,0,0,0,time.Local) },21},
		{Application{DOB:time.Date(2004,1,1,0,0,0,0,time.Local) },20},
		{Application{DOB:time.Date(2005,1,1,0,0,0,0,time.Local) },19},
		{Application{DOB:time.Date(2006,1,1,0,0,0,0,time.Local) },18},
		{Application{DOB:time.Date(2007,1,1,0,0,0,0,time.Local) },17},
		{Application{DOB:time.Date(2008,1,1,0,0,0,0,time.Local) },16},
		{Application{DOB:time.Date(2009,1,1,0,0,0,0,time.Local) },15},

		
	}
	for _,test:=range testCases{
		got:=test.application.CalculateAge()
		wanted:=test.result
		if got!=wanted{
			t.Errorf("Actual : %d Wanted : %d",got ,wanted)
		}
	}
}


func TestApplicationPost(t *testing.T){
	application:=Application{
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
	expected:=PostApplication{
		FatherFullName: "Asif Khan Baloch",
		FullName: "Numeez Khan Baloch",
		Age: 23,
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
	
	actual:=application.NewApplicationPost()
	id:=actual.Id
	expected.Id = id
	if actual!=expected{
		t.Errorf("Function is not returning the desired output")
	}

}