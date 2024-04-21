package app

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/zsandibe/postgres-pro-task/config"
	"github.com/zsandibe/postgres-pro-task/internal/storage"
	"github.com/zsandibe/postgres-pro-task/pkg/logger"
)

func Start() error {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("config.NewConfig: %v", err)
	}
	logger.Info("Config loaded successfully")

	db, err := storage.NewPostgresDB(cfg)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("storage.NewPostgresDB: %v", err)
	}
	defer db.Close()
	logger.Info("Database loaded successfully")

	if err = db.MigrateUp(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Error(err)
		return fmt.Errorf("storage.MigrateUp: %v", err)
	}

	return nil
}
