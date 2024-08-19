package main

import (
	"bytes"
	"context"
	"encoding/json"

	"errors"
	"time"

	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePostApplication(t *testing.T) {
	server := makeServer()
	request := makePostApplicationRequest()
	response := httptest.NewRecorder()
	AssertPostError(t, server, request, response)

	t.Run("Passing wrong body as a post", func(t *testing.T) {
		body, err := json.Marshal(struct{}{})
		if err != nil {
			log.Fatal(err)
		}
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		err = server.handlePostApplication(response, request)
		if err == nil {
			t.Error("Expecting an error but did not get it")
		}
	})

	t.Run("Making an application", func(t *testing.T) {
		application := Application{
			FirstName:        "Numeez",
			LastName:         "Baloch",
			MiddleName:       "Khan",
			FatherFirstName:  "Asif",
			FatherMiddleName: "Khan",
			FatherLastName:   "Baloch",
			Gender:           "male",
			City:             "Ahmedabad",
			StateOfDomicile:  "Gujarat",
			District:         "Ahmedabad",
			State:            "Gujarat",
			Village:          "Modasa",
			DOB:              time.Date(2001, time.September, 24, 0, 0, 0, 0, time.Local),
			HomeDistrict:     "Ahmedabad",
			HouseNo:          "1",
			Pincode:          380051,
			YearOfPassing:    "2023",
			Address:          "1 sterling silver heights",
		}
		body, err := json.Marshal(application)
		if err != nil {
			t.Fatalf("Failed to marshal application: %v", err)
		}
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		request.Header.Set("Content-Type", "application/json")

		err = server.handlePostApplication(response, request)
		if err != nil {
			t.Error("Error occured : ", err)
		}

		if response.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, response.Code)
		}

	})

}

func TestHandleDeleteApplication(t *testing.T) {
	server := makeServer()
	t.Run("Using any other method then POST", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		response := httptest.NewRecorder()
		err := server.handleDeleteApplicationById(response, request)
		expected := errors.New("HTTP METHOD should be POST is only allowed")
		if err.Error() != expected.Error() {
			t.Errorf("Expected error %s Actual error %s", expected.Error(), err.Error())
		}
	})
	t.Run("Deleting application with wrong or no Id", func(t *testing.T) {
		cases := []RequestId{{}, {"ffnsdlad1234"}}
		for _, c := range cases {
			body, err := json.Marshal(c)
			if err != nil {
				log.Fatal(err)
			}
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			response := httptest.NewRecorder()
			err = server.handleDeleteApplicationById(response, request)
			if err == nil {
				t.Error("Expected an error but did not get it")
			}
		}
	})
	t.Run("Deleting an application", func(t *testing.T) {
		var applications []PostApplication
		getRequest := httptest.NewRequest(http.MethodPost, "/", nil)
		getResponse := httptest.NewRecorder()
		err := server.handleFetchAllApplication(getResponse, getRequest)
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(getResponse.Body).Decode(&applications)
		id := applications[0].Id
		requestId := RequestId{Id: id}
		body, err := json.Marshal(requestId)
		if err != nil {
			log.Fatal(err)
		}
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		server.handleDeleteApplicationById(response, request)
		if response.Code != http.StatusOK {
			t.Errorf("Expected status code : %d Actual status code : %d ", http.StatusOK, response.Code)
		}
	})
}

func TestHandleApplicationById(t *testing.T) {
	server := makeServer()
	t.Run("Using any other method then post", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		err := server.handleDeleteApplicationById(response, request)
		expected := errors.New("HTTP METHOD should be POST is only allowed")
		if err.Error() != expected.Error() {
			t.Errorf("Expected error %s Actual error %s", expected.Error(), err.Error())
		}
	})
	t.Run("Not passing and id in query parameters", func(t *testing.T) {
		body, err := json.Marshal(struct{}{})
		if err != nil {
			log.Fatal(err)
		}
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		err = server.handleGetApplicationById(response, request)
		if err == nil {
			t.Error("Expected an error but did not get it")
		}
	})
	t.Run("Fetching a single application", func(t *testing.T) {
		id := RequestId{
			Id: "e6cf3a14-70f0-4cff-8bcf-b5dd284306bf",
		}
		body, err := json.Marshal(id)
		if err != nil {
			log.Fatal(err)
		}
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		err = server.handleGetApplicationById(response, request)
		if err != nil {
			t.Error(err)
		}
		var application PostApplication
		json.NewDecoder(response.Body).Decode(&application)
		got := application.Id
		if got != id.Id {
			t.Errorf("Expected : %s Got : %s", id.Id, got)
		}
	})
}

func TestHandleUpdateApplication(t *testing.T) {
	server := makeServer()
	t.Run("using another method then POST", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		err := server.handleUpdateApplication(response, request)
		expected := errors.New("HTTP METHOD should be POST is only")
		if err.Error() != expected.Error() {
			t.Errorf("Expected error %s Actual error %s", expected.Error(), err.Error())
		}
	})
	t.Run("sending update request without passing id in query parameters", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()
		err := server.handleUpdateApplication(response, request)
		expected := errors.New("id of the document to be updated is not provided")
		if err.Error() != expected.Error() {
			t.Errorf("Expected : %s Got : %s", expected.Error(), err.Error())
		}
	})
	t.Run("Updating fields which does not exists", func(t *testing.T) {
		data := Data{
			"college": "Lj",
		}
		body := makeBytes(data)
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		err := server.handleUpdateApplication(response, request)
		if err == nil {
			t.Error("Expected error but it did not give an error")
		}
	})
	t.Run("passig a wrong id in request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()
		err := server.handleUpdateApplication(response, request)
		if err == nil {
			t.Error("Expected error but it did not give an error")
		}

	})
	t.Run("Updating a field in an application", func(t *testing.T) {
		data := Data{
			"city": "Brampton",
		}
		body := makeBytes(data)
		request := httptest.NewRequest(http.MethodPost, "/?id=e6cf3a14-70f0-4cff-8bcf-b5dd284306bf", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		err := server.handleUpdateApplication(response, request)
		if err != nil {
			t.Error(err)
		}
		if response.Code != http.StatusOK {
			t.Errorf("Expected code : %d Got : %d", http.StatusOK, response.Code)
		}
	})
}

func TestHandleFetchAll(t *testing.T) {
	server := makeServer()
	t.Run("Using any other method then post", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		err := server.handleFetchAllApplication(response, request)
		expected := errors.New("HTTP METHOD POST is only allowed")
		if err.Error() != expected.Error() {
			t.Errorf("Expected error %s Actual error %s", expected.Error(), err.Error())
		}
	})
	t.Run("Passing wrong filters", func(t *testing.T) {
		filters := Data{
			"college": "Harvard",
		}
		body := makeBytes(filters)
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		err := server.handleFetchAllApplication(response, request)
		if err == nil {
			t.Error("Expected an error but did not get it")
		}
	})
	t.Run("Getting all applications", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()
		err := server.handleFetchAllApplication(response, request)
		if err != nil {
			t.Error(err)
		}
		if response.Code != http.StatusOK {
			t.Errorf("Expected : %d Wanted : %d", http.StatusOK, response.Code)
		}

	})
	t.Run("Getting all applications by applying filters", func(t *testing.T) {
		data := Data{
			"city": "Ahmedabad",
		}
		body:= makeBytes(data)
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		err := server.handleFetchAllApplication(response, request)
		if err != nil {
			t.Error(err)
		}
		if response.Code != http.StatusOK {
			t.Errorf("Expected : %d Wanted : %d", http.StatusOK, response.Code)
		}
	})

}

func AssertPostError(t testing.TB, server *Server, request *http.Request, response *httptest.ResponseRecorder) {
	t.Helper()
	err := server.handlePostApplication(response, request)
	expected := errors.New("HTTP METHOD POST is only")
	if err.Error() != expected.Error() {
		t.Errorf("Expected error %s Actual error %s", expected.Error(), err.Error())
	}
}

func makePostApplicationRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/makeapplication", nil)
}
func makeServer() *Server {
	store, err := NewMongoStore(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	server := NewApiServer(":3000", store)
	return server
}

func makeBytes(v any) []byte {
	body, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return body
}
