package log

import (
	"context"
	"log-service-proto/domain/entity"
)

type Log interface {
	Store(context.Context, *entity.Log) error
	Connect(context.Context) error
}
