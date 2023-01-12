package rpc

import (
	"log"
	"net"

	"github.com/maan19/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/maan19/hex/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
