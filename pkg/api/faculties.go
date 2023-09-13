package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webEngineering/pkg/helpers"
	"webEngineering/pkg/models"
)

func (api *API) FacultiesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uidStr := r.URL.Query().Get("uid")
		if uidStr == "" {
			http.Error(w, "University not specified", http.StatusBadRequest)
			return
		}
		uid, err := strconv.Atoi(uidStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		faculties, err := api.db.GetFacultiesByUID(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(faculties)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var faculty models.Faculty
		err := json.NewDecoder(r.Body).Decode(&faculty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !helpers.CheckFacultySafety(faculty) {
			http.Error(w, "Faculty not safe", http.StatusBadRequest)
			return
		}
		_, err = api.db.GetFacultiyByExam(faculty.ExamDate, faculty.ExamAud)
		if err == nil {
			http.Error(w, "Faculty with given exam info already exists", http.StatusBadRequest)
			return
		}
		id, err := api.db.PostFaculty(faculty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPatch:
		var faculty models.Faculty
		err := json.NewDecoder(r.Body).Decode(&faculty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !helpers.CheckFacultySafety(faculty) {
			http.Error(w, "Faculty not safe", http.StatusBadRequest)
			return
		}
		err = api.db.PatchFaculty(faculty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		var faculty models.Faculty
		err := json.NewDecoder(r.Body).Decode(&faculty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = api.db.DeleteFaculty(faculty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
