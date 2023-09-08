package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) PingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		str := r.URL.Query().Get("str")
		if str == "" {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:   "str",
			Value:  str,
			MaxAge: 20,
		})
		err := json.NewEncoder(w).Encode(str)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
