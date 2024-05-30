package grpc

import (
	logv1 "github.com/mahdimehrabi/m1-log-proto/gen/go/log/v1"
	"google.golang.org/grpc"
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
	logger.Info("running grpc server ⛴")
	env := godotenv.NewEnv()
	env.Load()

	logRepo := pgx.NewLogRepository(env)
	loggerService := logService.NewService(logger, logRepo)

	lis, err := net.Listen("tcp", env.ServerAddr)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	logServer := server.NewLogServer(logger, loggerService)
	logv1.RegisterLogServiceServer(grpcServer, logServer)
	err = grpcServer.Serve(lis)
	log.Fatal(err)
}
