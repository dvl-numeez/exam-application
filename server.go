package main

import (
	"log"
	"net/http"
)




type Server struct{
	listenAddr string
	store Storage
}


func NewApiServer(listenAddr string, store Storage)*Server{
	return &Server{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *Server)Start(){
	router:= http.NewServeMux()
	router.Handle("/health",http.HandlerFunc(health))
	err:=http.ListenAndServe(s.listenAddr,router)
	if err!=nil{
		log.Fatal("Unable to start the server due to error : ",err)
	}
	

}

func health(w http.ResponseWriter, r * http.Request){
	w.WriteHeader(http.StatusOK)
	
}