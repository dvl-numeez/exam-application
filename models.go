package main

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

const male = "male"
const female = "female"

type RequestId struct {
	Id string `json:"id"`
}

type Application struct {
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	MiddleName       string    `json:"middleName"`
	Gender           string    `json:"gender"`
	HomeDistrict     string    `json:"homeDistrict"`
	DOB              time.Time `json:"dob"`
	FatherFirstName  string    `json:"fatherFirstName"`
	FatherLastName   string    `json:"fatherLastName"`
	FatherMiddleName string    `json:"fatherMiddleName"`
	BoardName        string    `json:"boardName"`
	StateOfDomicile  string    `json:"stateOfDomicile"`
	YearOfPassing    string    `json:"yearOfPassing"`
	RollNumber       string    `json:"rollNumber"`
	Address          string    `json:"address"`
	State            string    `json:"state"`
	District         string    `json:"district"`
	City             string    `json:"city"`
	Pincode          int       `json:"pincode"`
	HouseNo          string    `json:"houseNo"`
	Village          string    `json:"village"`
}

type PostApplication struct {
	Id               string `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	MiddleName       string `json:"middleName"`
	FatherFirstName  string `json:"fatherFirstName"`
	FatherLastName   string `json:"fatherLastName"`
	FatherMiddleName string `json:"fatherMiddleName"`
	Age             int       `json:"age"`
	FullName        string    `json:"fullName"`
	FatherFullName  string    `json:"fatherFullName"`
	BoardName       string    `json:"boardName"`
	StateOfDomicile string    `json:"stateOfDomicile"`
	YearOfPassing   string    `json:"yearOfPassing"`
	RollNumber      string    `json:"rollNumber"`
	Address         string    `json:"address"`
	State           string    `json:"state"`
	District        string    `json:"district"`
	City            string    `json:"city"`
	Pincode         int       `json:"pincode"`
	HouseNo         string    `json:"houseNo"`
	Village         string    `json:"village"`
	Gender          string    `json:"gender"`
	HomeDistrict    string    `json:"homeDistrict"`
	DOB             time.Time `json:"dob"`
}

func (a *Application) ValidateGender() bool {
	gender := strings.ToLower(a.Gender)
	if gender == male || gender == female {
		return true
	}
	return false
}
func (a *Application) CalculateAge() int {
	presentYear := time.Now().Year()
	age := presentYear - a.DOB.Year()
	return age

}

func (a *Application) makeFullName(firstName, middleName, lastName string) string {
	return firstName + " " + middleName + " " + lastName
}

func (a *Application) NewApplicationPost() PostApplication {
	fullName := a.makeFullName(a.FirstName, a.MiddleName, a.LastName)
	fatherFullName := a.makeFullName(a.FatherFirstName, a.FatherMiddleName, a.FatherLastName)
	age := a.CalculateAge()
	return PostApplication{
		Id:              uuid.New().String(),
		Age:             age,
		FullName:        fullName,
		FatherFullName:  fatherFullName,
		District:        a.District,
		HomeDistrict:    a.HomeDistrict,
		HouseNo:         a.HouseNo,
		City:            a.City,
		DOB:             a.DOB,
		Address:         a.Address,
		BoardName:       a.BoardName,
		StateOfDomicile: a.StateOfDomicile,
		State:           a.State,
		Village:         a.Village,
		YearOfPassing:   a.YearOfPassing,
		RollNumber:      a.RollNumber,
		Pincode:         a.Pincode,
		Gender:          a.Gender,
		FirstName: a.FirstName,
		LastName: a.LastName,
		MiddleName: a.MiddleName,
		FatherFirstName: a.FatherFirstName,
		FatherLastName: a.FatherLastName,
		FatherMiddleName: a.FatherMiddleName,
	}
}
