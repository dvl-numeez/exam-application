package main

import (
	"encoding/json"
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

	}
	return nil
}

func(s *Server)handlePostApplication(w http.ResponseWriter, r *http.Request)error{
		var requestBody Application
		json.NewDecoder(r.Body).Decode(&requestBody)
		err:=s.store.InsertApplication(r.Context(),&requestBody)
		if err!=nil{
			return err
		}
		WriteJson(w,http.StatusCreated,map[string]string{"message":"application created"})
		return nil
}
func(s *Server)handleFetchAllApplication(w http.ResponseWriter, r *http.Request)error{
	applications,err:=s.store.FetchAll(r.Context())
	if err!=nil{
		return err
	}
	WriteJson(w,http.StatusOK,applications)
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