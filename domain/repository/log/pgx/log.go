package pgx

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log-service-proto/domain/entity"
	"log-service-proto/infrastructure/godotenv"
)

type LogRepository struct {
	env  *godotenv.Env
	conn *pgx.Conn
}

func NewLogRepository(env *godotenv.Env) *LogRepository {
	lr := &LogRepository{
		env: env,
	}
	return lr
}

func (r LogRepository) Connect(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, r.env.PGSQLConnection)
	if err != nil {
		return err
	}
	r.conn = conn
	return nil
}

func (r LogRepository) Store(ctx context.Context, log *entity.Log) error {
	if _, err := r.conn.Exec(ctx, `INSERT INTO logs (created_at,error) VALUES($1,$2)  `,
		log.CreatedAt, log.Error); err != nil {
		return err
	}
	return nil
}
