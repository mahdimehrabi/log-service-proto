package server

import (
	"context"
	"errors"
	logv1 "github.com/mahdimehrabi/m1-log-proto/gen/go/log/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"m1-log-service/domain/entity"
	"m1-log-service/domain/service/log"
	logger "m1-log-service/infrastructure/log"
)

type LogServer struct {
	logger     logger.Logger
	logService *log.Service

	logv1.UnimplementedLogServiceServer
}

func NewLogServer(logger logger.Logger, logService *log.Service) *LogServer {
	return &LogServer{logger: logger, logService: logService}
}

func (l LogServer) StoreLog(server logv1.LogService_StoreLogServer) error {
	chErr := make(chan error)
	go func() {
		for {
			logErr, err := server.Recv()
			if err != nil {
				chErr <- err
				return
			}
			log := entity.NewLog(logErr.Error)
			if err := l.logService.Store(context.Background(), log); err != nil {
				chErr <- err
			}
		}
	}()

	<-chErr
	if err := server.SendAndClose(&logv1.Empty{}); err != nil {
		if !errors.Is(err, context.Canceled) {
			l.logger.Error(err)
		}
	}

	return status.Errorf(codes.Internal, "internal server error")
}
