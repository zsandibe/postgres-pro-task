package main

import (
	"github.com/zsandibe/postgres-pro-task/internal/app"
	"github.com/zsandibe/postgres-pro-task/pkg/logger"
)

func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
