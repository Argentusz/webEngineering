package storage

import (
	"context"
	"webEngineering/pkg/models"
)

func (s *Storage) GetStudent(id int) (models.Student, error) {
	var student models.Student
	err := s.pool.QueryRow(context.Background(),
		`SELECT id, name, exam FROM students WHERE id = $1`, id).
		Scan(&student.ID, &student.Name, &student.Exam)
	return student, err
}

func (s *Storage) GetAllStudents() ([]models.Student, error) {
	var res []models.Student
	rows, err := s.pool.Query(context.Background(),
		`SELECT id, name, exam FROM students`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.ID, &student.Name, &student.Exam)
		if err != nil {
			return nil, err
		}
		res = append(res, student)
	}
	return res, nil
}

func (s *Storage) PostStudent(student models.Student) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(),
		`INSERT INTO students(name, exam) VALUES ($1, $2) RETURNING id`, student.Name, student.Exam).Scan(&id)
	return id, err
}

func (s *Storage) PatchStudent(student models.Student) error {
	_, err := s.pool.Query(context.Background(),
		`UPDATE students SET name = $2, exam = $3 WHERE id = $1`, student.ID, student.Name, student.Exam)
	return err
}

func (s *Storage) DeleteStudent(student models.Student) error {
	_ = s.pool.QueryRow(context.Background(),
		`DELETE FROM faculties_to_students WHERE studentid = $1`, student.ID).Scan()
	_ = s.pool.QueryRow(context.Background(),
		`DELETE FROM students WHERE id = $1`, student.ID).Scan()
	return nil
}
