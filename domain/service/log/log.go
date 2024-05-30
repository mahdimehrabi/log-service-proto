package log

import (
	"context"
	"m1-log-service/domain/entity"
	logRepo "m1-log-service/domain/repository/log"
	loggerInfra "m1-log-service/infrastructure/log"
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
