package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"webEngineering/pkg/helpers"
	"webEngineering/pkg/models"
)

func (api *API) StudentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			students, err := api.db.GetAllStudents()
			if err != nil {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(w).Encode(students)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			student, err := api.db.GetStudent(id)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(w).Encode(student)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	case http.MethodPost:
		var student models.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !helpers.CheckStudentSafety(student) {
			log.Println("Student", student, "not safe")
			http.Error(w, "Student not safe", http.StatusBadRequest)
			return
		}
		id, err := api.db.PostStudent(student)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(id)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPatch:
		var student models.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = api.db.PatchStudent(student)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		var student models.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = api.db.DeleteStudent(student)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
