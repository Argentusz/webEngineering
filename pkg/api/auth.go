package api

import (
	"encoding/json"
	"net/http"
	"webEngineering/pkg/helpers"
	"webEngineering/pkg/models"
)

func (api *API) AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		login, password, ok := r.BasicAuth()
		if !ok {
			unauthorized(w)
			return
		}
		dbPassword, err := api.db.GetPassword(login)
		if err != nil || !comparePasswords(password, dbPassword) {
			unauthorized(w)
			return
		}
		university, err := api.db.GetUniversityByLogin(login)
		if err != nil {
			unauthorized(w)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "lang",
			Value: university.Lang,
		})
		err = json.NewEncoder(w).Encode(university.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		var university models.University
		err := json.NewDecoder(r.Body).Decode(&university)
		if err != nil {
			unauthorized(w)
			return
		}
		if !helpers.CheckUniversitySafety(university) {
			unauthorized(w)
			return
		}
		// Check if user with this login already exists
		_, err = api.db.GetUniversityByLogin(university.Login)
		if err == nil {
			unauthorized(w)
			return
		}
		newUId, err := api.db.NewUniversity(university)
		if err != nil {
			unauthorized(w)
			return
		}
		err = json.NewEncoder(w).Encode(newUId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
