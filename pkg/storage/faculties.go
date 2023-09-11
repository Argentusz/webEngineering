package storage

import (
	"context"
	"webEngineering/pkg/models"
)

func (s *Storage) GetFacultyByUID(uid int) (models.Faculty, error) {
	var faculty models.Faculty
	err := s.pool.QueryRow(context.Background(),
		`SELECT id, universityID, name, exam_date, exam_aud FROM faculties WHERE universityID = $1`, uid).
		Scan(&faculty.ID, &faculty.UniversityID, &faculty.Name, &faculty.ExamDate, &faculty.ExamAud)
	return faculty, err
}

func (s *Storage) PostFaculty(faculty models.Faculty) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(),
		`INSERT INTO faculties(universityID, name, exam_date, exam_aud) VALUES ($1, $2, $3, $4) RETURNING id`).
		Scan(&id)
	return id, err
}

func (s *Storage) PatchFaculty(faculty models.Faculty) error {
	_, err := s.pool.Query(context.Background(),
		`UPDATE faculties SET name = $2, universityid = $3, exam_date = $4, exam_aud = $5 WHERE id = $1`,
		faculty.ID, faculty.Name, faculty.UniversityID, faculty.ExamDate, faculty.ExamAud)
	return err
}

func (s *Storage) DeleteFaculty(faculty models.Faculty) error {
	_ = s.pool.QueryRow(context.Background(),
		`DELETE FROM faculties_to_students WHERE courseid = $1`, faculty.ID)
	_ = s.pool.QueryRow(context.Background(),
		`DELETE FROM faculties WHERE id = $1`, faculty.ID)
	return nil
}
