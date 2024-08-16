package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type Server struct {
	listenAddr string
	store      Storage
}

type ApiError struct {
	Error string `json:"error"`
}

func NewApiServer(listenAddr string, store Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() {
	router := http.NewServeMux()
	router.Handle("/application", makeHttpHandleFunc(s.handleGetApplicationById))
	router.Handle("/getallapplications", makeHttpHandleFunc(s.handleFetchAllApplication))
	router.Handle("/deleteapplication", makeHttpHandleFunc(s.handleDeleteApplicationById))
	router.Handle("/updateapplication", makeHttpHandleFunc(s.handleUpdateApplication))
	router.Handle("/makeapplication", makeHttpHandleFunc(s.handlePostApplication))
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatal("Unable to start the server due to error : ", err)
	}

}

func (s *Server) handlePostApplication(w http.ResponseWriter, r *http.Request) error {
	requestBody:= Application{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return err
	}
	err = s.store.InsertApplication(r.Context(), &requestBody)
	if err != nil {
		return err
	}
	err = WriteJson(w, http.StatusCreated, map[string]string{"message": "application created"})
	if err != nil {
		return err
	}
	return nil
}
func (s *Server) handleFetchAllApplication(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("HTTP METHOD POST is only")
	}
	var filters map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&filters)
	if err != nil {
		filters = map[string]interface{}{}
	}
	applications, err := s.store.FetchAll(r.Context(), filters)
	if err != nil {
		return err
	}
	err = WriteJson(w, http.StatusOK, applications)
	if err != nil {
		return err
	}
	return nil
}
func (s *Server) handleDeleteApplicationById(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("HTTP METHOD POST is only")
	}

	requestId :=RequestId{}
	err := json.NewDecoder(r.Body).Decode(&requestId)
	if err != nil {
		return err
	}
	err = s.store.Delete(r.Context(), requestId.Id)
	if err != nil {
		return err
	}
	err = WriteJson(w, http.StatusOK, map[string]string{"message": "Application successfully deleted"})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) handleGetApplicationById(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("HTTP METHOD POST is only")
	}
	requestId := RequestId{}
	err := json.NewDecoder(r.Body).Decode(&requestId)
	if err != nil {
		return err
	}
	application, err := s.store.GetApplicationById(r.Context(), requestId.Id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = WriteJson(w, http.StatusOK, application)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) handleUpdateApplication(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("HTTP METHOD POST is only")
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		return errors.New("id of the document to be updated is not provided")
	}
	var filters map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&filters)
	if err != nil {
		return err
	}
	err = s.store.UpdateApplication(r.Context(), filters, id)
	if err != nil {
		return err
	}
	WriteJson(w, http.StatusOK, map[string]string{"message": "application data updated"})
	return nil
}

func makeHttpHandleFunc(function apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := function(w, r)
		if err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
