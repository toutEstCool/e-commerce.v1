package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	user "github.com/toutEstCool/e-commerce.v1/cmd/service"
)

type APIServer struct {
	addres string
	db     *sql.DB
}

func NewAPIServer(addres string, db *sql.DB) *APIServer {
	return &APIServer{
		addres: addres,
		db:     db,
	}
}

func (server *APIServer) Run() error {
	route := mux.NewRouter()
	subroute := route.PathPrefix("/api/v1").Subrouter()

	userService := user.NewHandler()
	userService.RegisterRoutes(subroute)

	log.Println("Listening on", server.addres)
	return http.ListenAndServe(server.addres, route)
}
