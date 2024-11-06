package application

import (
	"github.com/zuekveer/BWA/internal/databases"
	"github.com/zuekveer/BWA/internal/repository"
	"github.com/zuekveer/BWA/internal/service"
)

type eventDomain struct {
	events *service.EventService
}

func buildEventDomain(db *databases.DB) eventDomain {
	repo := repository.NewRepository(db)
	userService := service.NewEventService(repo)
	return eventDomain{userService}
}
