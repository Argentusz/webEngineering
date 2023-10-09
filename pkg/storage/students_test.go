package storage

import (
	"fmt"
	"testing"
	"webEngineering/pkg/config"
	"webEngineering/pkg/models"
)

func TestStorage_Students(t *testing.T) {
	var cfg config.Config
	cfg.Setup()
	s, err := New(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbName,
	))
	if err != nil {
		t.Fatal("Could not connect to database")
	}
	student := models.Student{Name: "Test Student", Exam: 5}

	newSid, err := s.PostStudent(student)
	if newSid == 0 || err != nil {
		t.Fatal("Could not post student")
	}

	gotStudent, err := s.GetStudent(newSid)
	if err != nil {
		t.Error("Could not get student")
	} else if gotStudent.Name != student.Name || gotStudent.Exam != student.Exam {
		t.Error("Mismatch in posted and got student", student, gotStudent)
	}

	students, err := s.GetAllStudents()
	if err != nil {
		t.Error("Could not get all students:", err.Error())
	} else if len(students) == 0 {
		t.Error("Could not get all students:", students)
	}

	student.ID = newSid
	student.Name = "Changed Test Student"
	err = s.PatchStudent(student)
	if err != nil {
		t.Error("Could not patch student:", err.Error())
	}

	err = s.DeleteStudent(student)
	if err != nil {
		t.Error("Could not delete student", err.Error())
	}
}
