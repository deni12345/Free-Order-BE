package main

import (
	"github/lambda-microservice/api"
	"github/lambda-microservice/api/middleware"
	"github/lambda-microservice/custom"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	s := api.NewServer()

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(middleware.ContentTypeJsonMiddleware)
	private(r, s)
	http.Handle("/", r)
	custom.PrintLogo()
	http.ListenAndServe(":8080", nil)
}

// func newServer(l logic.Logic) *mux.Route {
// 	r := mux.NewRouter()
// 	api := r.PathPrefix("/api/").Subrouter()
// 	public(api, l)
// 	return r
// }

func private(r *mux.Router, s api.Server) {
	pri := r.PathPrefix("/private").Subrouter()
	pri.HandleFunc("/create-user", s.SignUp).Methods("GET")
}
