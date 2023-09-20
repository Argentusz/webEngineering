package api

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func (api *API) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api") || r.URL.Path == "/auth" || r.URL.Path == "/ping" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Content-Type", "application/json")
			if r.Method == http.MethodOptions {
				return
			}
		}

		if strings.Contains(r.URL.Path, "/api") {
			// HTTP Authorization
			login, password, ok := r.BasicAuth()
			if !ok {
				fmt.Println("Middleware error: could not parse auth token")
				unauthorized(w)
				return
			}
			dbPassword, err := api.db.GetPassword(login)
			if err != nil || !comparePasswords(password, dbPassword) {
				fmt.Printf("error while identification: dbPassword=\"%s\", login=\"%s\" password=\"%s\"",
					dbPassword, login, password)
				unauthorized(w)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("Notes-WWW-Authenticate", `Basic realm="api"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func comparePasswords(password, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	return err == nil
}
