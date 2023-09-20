package config

import (
	"fmt"
	"testing"
	"webEngineering/pkg/storage"
)

func TestBD(t *testing.T) {
	var cfg Config
	cfg.Setup()
	_, err := storage.New(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbName,
	))
	if err != nil {
		t.Error(err.Error())
	}
}
