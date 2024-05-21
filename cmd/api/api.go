package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/steevedevops/first-go-api/cmd/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// inicializa a structura de APIserver em programacao oriendado a objeto
// essa parte seria o contrutor da clase... em dart seria o initialize

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// cria um metodo da estrutura
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
