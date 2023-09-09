package storage

import "context"

func (s *Storage) GetPassword(login string) (string, error) {
	var password string
	err := s.pool.QueryRow(context.Background(),
		`SELECT password FROM universities WHERE login = $1`, login).Scan(&password)
	return password, err
}
