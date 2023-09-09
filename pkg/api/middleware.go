package api

import (
	"net/http"
	"strings"
)

func (api *API) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api") {
			login, password, ok := r.BasicAuth()
			if !ok {
				unauthorized(w)
				return
			}
			dbPassword, err := api.db.GetPassword(login)
			if err != nil {
				unauthorized(w)
				return
			}
			if !comparePasswords(password, dbPassword) {
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
	return password == dbPassword
	//err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	//return err == nil
}
