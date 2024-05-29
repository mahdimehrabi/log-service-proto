package log

import (
	"context"
	"log-service-proto/domain/entity"
	logRepo "log-service-proto/domain/repository/log"
	loggerInfra "log-service-proto/infrastructure/log"
)

type Service struct {
	logRepository logRepo.Log
	logger        loggerInfra.Logger
}

func NewService(logger loggerInfra.Logger, logRep logRepo.Log) *Service {
	return &Service{
		logRepository: logRep,
		logger:        logger,
	}
}

func (s Service) Store(ctx context.Context, logEnt *entity.Log) error {
	err := s.logRepository.Store(ctx, logEnt)
	if err != nil {
		s.logger.Error(err)
	}
	return err
}
