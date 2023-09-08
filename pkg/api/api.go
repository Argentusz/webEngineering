package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"webEngineering/pkg/config"
	"webEngineering/pkg/storage"
)

type API struct {
	router *mux.Router
	db     *storage.Storage
	Server
}

type Server struct {
	address string
}

func New(cfg config.Config) (*API, error) {
	s, err := storage.New(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbName,
	))
	if err != nil {
		return nil, err
	}
	router := mux.NewRouter()

	return &API{
		router: router,
		db:     s,
		Server: Server{
			address: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		},
	}, nil
}

func (api *API) Fill() {
	// TODO: api.router.HandleFunc("api/...", <handler>).Methods(http.MethodGet, ... )
	api.router.HandleFunc("/api/ping", api.PingHandler).Methods(http.MethodGet)
}

func (api *API) Serve() error {
	return http.ListenAndServe(api.address, api.router)
}
