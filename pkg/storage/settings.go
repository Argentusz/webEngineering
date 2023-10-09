package storage

import (
	"context"
	"webEngineering/pkg/models"
)

func (s *Storage) PatchSettings(settings models.Settings) error {
	_, err := s.pool.Query(context.Background(),
		`UPDATE universities SET lang = $2 WHERE id = $1`, settings.Uid, settings.Lang)
	return err
}
