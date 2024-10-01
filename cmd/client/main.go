package main

import (
	"flag"

	pb "github.com/echelon/api/grpc"
	"github.com/echelon/config"
	"github.com/echelon/internal/client/service"
	z "github.com/echelon/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf, err := config.LoadConfig("config/client.yaml")
	if err != nil {
		z.Logger.Errorf("load config: %v", err)
	}

	asyncFlag := flag.Bool("async", false, "Download thumbnails asynchronously")
	flag.Parse()

	videoURLs := flag.Args()
	if len(videoURLs) == 0 {
		z.Logger.Errorf("No video URLs provided")
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(conf.GrpcPort, opts...)
	if err != nil {
		z.Logger.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer func() {
		er := conn.Close()
		if er != nil {
			z.Logger.Errorf("closing connection: %v", er)
		}
	}()

	client := pb.NewThumbnailServiceClient(conn)

	if *asyncFlag {
		service.DownloadThumbnailsAsync(client, videoURLs)
	} else {
		service.DownloadThumbnailsSync(client, videoURLs)
	}
}
