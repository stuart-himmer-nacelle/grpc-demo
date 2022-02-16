package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func newServer() *logService {
	s := &logService{}
	return s
}

func main() {
	flag.Parse()
	fmt.Println("port:", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	RegisterLogServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

type logService struct {
	UnimplementedLogServiceServer
}

func (ls *logService) SaveLog(ctx context.Context, in *Log) (*OutLog, error) {
	log := OutLog{
		Severity:  in.GetSeverity(),
		Message:   in.GetMessage(),
		Timestamp: time.Now().GoString(),
	}
	return &log, nil
}
