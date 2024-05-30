package grpc

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	logv1 "github.com/mahdimehrabi/m1-log-proto/gen/go/log/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"log-service-proto/application/grpc/server"
	"log-service-proto/domain/repository/log/pgx"
	logService "log-service-proto/domain/service/log"
	"log-service-proto/infrastructure/godotenv"
	"log-service-proto/infrastructure/log/zerolog"
	"net"
)

func Boot() {
	logger := zerolog.NewLogger()
	env := godotenv.NewEnv()
	env.Load()

	conn, err := pgxpool.New(context.Background(), env.DATABASE_HOST)
	if err != nil {
		log.Fatal(err)
	}

	logRepo := pgx.NewLogRepository(env, conn)
	if err != nil {
		log.Fatal(err)
	}
	loggerService := logService.NewService(logger, logRepo)

	lis, err := net.Listen("tcp", env.ServerAddr)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	logServer := server.NewLogServer(logger, loggerService)
	logv1.RegisterLogServiceServer(grpcServer, logServer)

	reflection.Register(grpcServer)
	logger.Info(fmt.Sprintf("running grpc server on: %s ⛴", env.ServerAddr))
	err = grpcServer.Serve(lis)
	log.Fatal(err)
}
