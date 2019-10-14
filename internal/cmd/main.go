package main

import (
	"net"

	impl "github.com/emart.io/zbay/internal/impl/service"
	pb "github.com/emart.io/zbay/service/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	port = ":50051"
)

func main() {
	go serveFileUpload()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommoditiesServer(s, new(impl.CommoditiesImpl))
	pb.RegisterUsersServer(s, new(impl.UsersImpl))
	pb.RegisterMessagesServer(s, new(impl.MessageImpl))

	grpclog.Println("begin..." + port)
	if err := s.Serve(lis); err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}
