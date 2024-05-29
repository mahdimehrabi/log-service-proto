package log

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"log-service-proto/domain/entity"
	infraMock "log-service-proto/mock/infrastructure"
	repoMock "log-service-proto/mock/repository"
	"testing"
	"time"
)

func TestService_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name        string
		log         *entity.Log
		loggerMock  func() *infraMock.MockLog
		logRepoMock func() *repoMock.MockLog
		error       error
		ctx         context.Context
	}{
		{
			name: "success",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				return loggerInfra
			},
			logRepoMock: func() *repoMock.MockLog {
				repoLogMock := repoMock.NewMockLog(ctrl)
				repoLogMock.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)
				return repoLogMock
			},
			log: &entity.Log{
				Error:     "error",
				CreatedAt: uint64(time.Now().Unix()),
			},
			error: nil,
			ctx:   context.Background(),
		},
		{
			name: "RepoError",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			logRepoMock: func() *repoMock.MockLog {
				repoLogMock := repoMock.NewMockLog(ctrl)
				repoLogMock.EXPECT().Store(gomock.Any(), gomock.Any()).Return(err)
				return repoLogMock
			},
			log: &entity.Log{
				Error:     "error",
				CreatedAt: uint64(time.Now().Unix()),
			},
			error: err,
			ctx:   context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.logRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			err := service.Store(test.ctx, test.log)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}
}
