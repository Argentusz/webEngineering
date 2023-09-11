package storage

import (
	"context"
	"webEngineering/pkg/models"
)

func (s *Storage) GetStudentsByFacultyID(fid int) ([]models.Student, error) {
	var students []models.Student
	var sids []int32
	rows, err := s.pool.Query(context.Background(),
		`SELECT studentID FROM faculties_to_students WHERE courseID = $1`, fid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var sid int
		err = rows.Scan(&sid)
		if err != nil {
			return nil, err
		}
		sids = append(sids, int32(sid))
	}

	sRows, err := s.pool.Query(context.Background(),
		`SELECT id, name, exam FROM students WHERE id = ANY($1)`, sids)
	defer sRows.Close()

	for sRows.Next() {
		var student models.Student
		err = sRows.Scan(
			&student.ID,
			&student.Name,
			&student.Exam,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *Storage) PostFacultiesToStudents(fid, sid int) error {
	var id int
	err := s.pool.QueryRow(context.Background(),
		`INSERT INTO faculties_to_students(courseid, studentid) VALUES ($1, $2) RETURNING courseid`, fid, sid).Scan(&id)
	return err
}

func (s *Storage) DeleteFacultiesToStudents(fid, sid int) error {
	_ = s.pool.QueryRow(context.Background(),
		`DELETE FROM faculties_to_students WHERE courseID = $1 AND studentID = $2`, fid, sid)
	return nil
}
