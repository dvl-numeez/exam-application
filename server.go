package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)


type apiFunc func(w http.ResponseWriter,r *http.Request)error

type Server struct{
	listenAddr string
	store Storage
}

type ApiError struct{
	Error string `json:"error"`
}


func NewApiServer(listenAddr string, store Storage)*Server{
	return &Server{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *Server)Start(){
	router:= http.NewServeMux()
	router.Handle("/application",makeHttpHandleFunc(s.handleApplication))
	err:=http.ListenAndServe(s.listenAddr,router)
	if err!=nil{
		log.Fatal("Unable to start the server due to error : ",err)
	}
	

}

func(s *Server)handleApplication(w http.ResponseWriter, r *http.Request)error{
	switch r.Method{	
	case "POST":
		return s.handlePostApplication(w,r)
	case "GET":
		return s.handleFetchAllApplication(w,r)
	case "DELETE":
		return s.handleDeleteApplicationById(w,r)
	case "PATCH":
		return s.handleUpdateApplication(w,r)

	}
	return nil
}

func(s *Server)handlePostApplication(w http.ResponseWriter, r *http.Request)error{
		var requestBody Application
		err:=json.NewDecoder(r.Body).Decode(&requestBody)
		if err!=nil{
			return err
		}
		err=s.store.InsertApplication(r.Context(),&requestBody)
		if err!=nil{
			return err
		}
		WriteJson(w,http.StatusCreated,map[string]string{"message":"application created"})
		return nil
}
func(s *Server)handleFetchAllApplication(w http.ResponseWriter, r *http.Request)error{
	id:=r.URL.Query().Get("id")
	if id!=""{
		return s.handleDeleteApplicationById(w,r)
	}
	applications,err:=s.store.FetchAll(r.Context())
	if err!=nil{
		return err
	}
	WriteJson(w,http.StatusOK,applications)
	return nil
}
func(s *Server)handleDeleteApplicationById(w http.ResponseWriter, r *http.Request)error{
	var requestId RequestId 
	err:=json.NewDecoder(r.Body).Decode(&requestId)
	if err!=nil{
		return err
	}
	err=s.store.Delete(r.Context(),requestId.Id)
	if err!=nil{
		return err
	}
	WriteJson(w,http.StatusOK,map[string]string{"message":"Application successfully deleted"})
	return nil
}

func(s *Server)handleGetApplicationById(w http.ResponseWriter, r *http.Request)error{
	id:=r.URL.Query().Get("id")
	application,err:=s.store.GetApplicationById(r.Context(),id)
	if err!=nil{
		return err
	}
	WriteJson(w,http.StatusOK,application)
	return nil
}

func(s *Server)handleUpdateApplication(w http.ResponseWriter,r *http.Request)error{
	id:=r.URL.Query().Get("id")
	if id==""{
		return errors.New("id of the document to be updated is not provided")
	}
	var application map[string]interface{}
	err:=json.NewDecoder(r.Body).Decode(&application)
	if err!=nil{
		return err
	}
	err=s.store.UpdateApplication(r.Context(),application,id)
	if err!=nil{
		return nil
	}
	WriteJson(w,http.StatusOK,map[string]string{"message":"application data updated"})
	return nil
}



func makeHttpHandleFunc(function apiFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		err:=function(w,r)
		if err!=nil{
			WriteJson(w,http.StatusBadRequest,ApiError{Error: err.Error()})
		}
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}