package api

import (
	"encoding/json"
	"log"
	"net/http"
	"webEngineering/pkg/models"
)

func (api *API) SettingsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		var settings models.Settings
		err := json.NewDecoder(r.Body).Decode(&settings)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = api.db.PatchSettings(settings)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
