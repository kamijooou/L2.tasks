package api

import (
	"sched/internal/logger"
	"sched/internal/service"
	"sched/internal/storage/inmemory"
)

func InitializeServer() (*server, error) {
	storage := inmemory.NewStorage()
	calendarService := service.NewCalendarService(storage)
	logrusLogger := logger.NewLogger()
	apiServer := newServer(calendarService, storage, logrusLogger)
	return apiServer, nil
}
