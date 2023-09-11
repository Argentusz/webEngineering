package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webEngineering/pkg/models"
)

func (api *API) FacultiesToStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		facultyID := r.URL.Query().Get("id")
		if facultyID == "" {
			http.Error(w, "Faculty ID not specified", http.StatusBadRequest)
			return
		}
		fid, err := strconv.Atoi(facultyID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		students, err := api.db.GetStudentsByFacultyID(fid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(students)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var pair models.FacultyToStudent
		err := json.NewDecoder(r.Body).Decode(&pair)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if pair.Fid == 0 || pair.Sid == 0 {
			http.Error(w, "Faculty/Student ID not specified", http.StatusBadRequest)
			return
		}
		err = api.db.PostFacultiesToStudents(pair.Fid, pair.Sid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		var pair models.FacultyToStudent
		err := json.NewDecoder(r.Body).Decode(&pair)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if pair.Fid == 0 || pair.Sid == 0 {
			http.Error(w, "Faculty/Student ID not specified", http.StatusBadRequest)
			return
		}
		err = api.db.DeleteFacultiesToStudents(pair.Fid, pair.Sid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
