package log

import (
	"context"
	"m1-log-service/domain/entity"
)

type Log interface {
	Store(context.Context, *entity.Log) error
}
