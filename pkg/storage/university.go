package storage

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"webEngineering/pkg/models"
)

func (s *Storage) GetPassword(login string) (string, error) {
	var password string
	err := s.pool.QueryRow(context.Background(),
		`SELECT password FROM universities WHERE login = $1`, login).Scan(&password)
	return password, err
}

func (s *Storage) GetID(login string) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(),
		`SELECT id FROM universities WHERE login = $1`, login).Scan(&id)
	return id, err
}

func (s *Storage) NewUniversity(university models.University) (int, error) {
	var id int
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(university.Password), 10)
	err = s.pool.QueryRow(context.Background(),
		`INSERT INTO universities(name, login, password) VALUES ($1, $2, $3) RETURNING id`,
		university.Name, university.Login, string(hashedPassword)).Scan(&id)
	return id, err
}
