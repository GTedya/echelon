package main

import (
	"fmt"
	"net"

	pb "github.com/echelon/api/grpc"
	"github.com/echelon/config"
	"github.com/echelon/internal/server/repository"
	"github.com/echelon/internal/server/repository/sqlite"
	"github.com/echelon/internal/server/service"
	z "github.com/echelon/pkg/logger"

	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadConfig("config/server.yaml")
	if err != nil {
		z.Logger.Fatalf("Failed to load config: %v", err)
	}

	repo, err := initDB(conf.DBPath)
	if err != nil {
		z.Logger.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		er := repo.Close()
		if er != nil {
			z.Logger.Errorf("repo closing: %v", er)
		}
	}()

	if err = runGrpcServer(conf.GrpcPort, repo); err != nil {
		z.Logger.Fatalf("Failed to start gRPC server: %v", err)
	}
}

func initDB(dbPath string) (*sqlite.LiteRepo, error) {
	repo := &sqlite.LiteRepo{}
	if err := repo.Open(dbPath); err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	return repo, nil
}

func runGrpcServer(port string, repo repository.ThumbNailRepository) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", port, err)
	}

	srv := service.NewCacheThumbnailService(repo)
	grpcServer := grpc.NewServer()
	thumbnailService := pb.NewThumbnailService(srv)
	pb.RegisterThumbnailServiceServer(grpcServer, thumbnailService)

	z.Logger.Debugf("Starting gRPC server on :%s", port)

	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}
