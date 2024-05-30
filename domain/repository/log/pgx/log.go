package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"m1-log-service/domain/entity"
	"m1-log-service/infrastructure/godotenv"
)

type LogRepository struct {
	env  *godotenv.Env
	conn *pgxpool.Pool
}

func NewLogRepository(env *godotenv.Env, conn *pgxpool.Pool) *LogRepository {
	lr := &LogRepository{
		env:  env,
		conn: conn,
	}
	return lr
}

func (r LogRepository) Store(ctx context.Context, log *entity.Log) error {
	if _, err := r.conn.Exec(ctx, `INSERT INTO logs (created_at,error) VALUES($1,$2)  `,
		log.CreatedAt, log.Error); err != nil {
		return err
	}
	return nil
}
