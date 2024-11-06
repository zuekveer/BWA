package application

import (
	"github.com/zuekveer/BWA/internal/databases"
	"github.com/zuekveer/BWA/internal/logger"
	"github.com/zuekveer/BWA/internal/repository"
	"github.com/zuekveer/BWA/internal/service"
)

type cronDomain struct {
	service *service.CronService
}

func buildCronDomain(db *databases.DB, l logger.Logger) cronDomain {
	repo := repository.NewRepository(db)
	cronService := service.NewCronService(repo, l)
	return cronDomain{cronService}
}
