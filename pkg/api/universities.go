package api

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"webEngineering/pkg/models"
)

func (api *API) UniversitiesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		var university models.University
		err := json.NewDecoder(r.Body).Decode(&university)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(university.Password), 10)
		university.Password = string(hashedPassword)
		err = api.db.PatchUniversity(university)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
