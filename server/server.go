package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiService struct {
	listenAddr string
	apiHandler *ApiHandler
}

func NewApiService(listenAddr string) *ApiService {
	s := &ApiService{
		listenAddr: listenAddr,
	}

	s.apiHandler = newApiHandler()
	return s
}

func (s *ApiService) Run() {
	fmt.Printf("server running on  port: %v\n", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, s.apiHandler.router))
}

func (s *ApiService) Use(middleware func(*ApiHandler)) {
	middleware(s.apiHandler)
}

type ApiHandler struct {
	router *mux.Router
}

func newApiHandler() *ApiHandler {
	return &ApiHandler{
		router: mux.NewRouter().StrictSlash(true),
	}
}

func (s *ApiHandler) Use(middleware func(*ApiHandler)) {
	middleware(s)
}

func (s *ApiHandler) Get(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f).Methods("GET")
}

func (s *ApiHandler) Post(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f).Methods("POST")
}

func (s *ApiHandler) Put(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f).Methods("PUT")
}

func (s *ApiHandler) Delete(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f).Methods("DELETE")
}
