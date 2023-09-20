package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
	"webEngineering/pkg/config"
	"webEngineering/pkg/storage"
)

type API struct {
	router *mux.Router
	db     *storage.Storage
	Server
}

type Server struct {
	address     string
	certificate string
	key         string
	frontend    string
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
			address:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			certificate: cfg.CertFile,
			key:         cfg.KeyFile,
			frontend:    cfg.FrontEnd,
		},
	}, nil
}

func (api *API) Fill() {
	_, err := os.Stat(api.frontend)
	for os.IsNotExist(err) {
		log.Println("Folder", api.frontend, "does not exist.")
		time.Sleep(2 * time.Second)
		_, err = os.Stat(api.frontend)
	}
	log.Println(api.frontend, "found")

	api.db.SetupDB()

	api.router.Use(api.Middleware)
	api.router.HandleFunc("/ping", api.PingHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/auth", api.AuthHandler).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/api/universities", api.UniversitiesHandler).Methods(http.MethodPatch, http.MethodOptions)
	api.router.HandleFunc("/api/students", api.StudentsHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions)
	api.router.HandleFunc("/api/faculties", api.FacultiesHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions)
	api.router.HandleFunc("/api/faculties_to_students", api.FacultiesToStudents).Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions)
	// Serve static assets directly.
	api.router.PathPrefix("/css").Handler(http.FileServer(http.Dir(api.frontend)))
	api.router.PathPrefix("/js").Handler(http.FileServer(http.Dir(api.frontend)))
	api.router.PathPrefix("/img").Handler(http.FileServer(http.Dir(api.frontend)))
	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	api.router.PathPrefix("/").HandlerFunc(IndexHandler(api.frontend + "/index.html"))
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

func (api *API) Serve() error {
	return http.ListenAndServeTLS(api.address, api.certificate, api.key, api.router)
}
