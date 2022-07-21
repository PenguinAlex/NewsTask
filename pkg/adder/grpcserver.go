package adder

import (
	api "awesomeProject"
	api2 "awesomeProject/pkg/api"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

//GRPCServer ...
type GRPCServer struct {
	api2.UnimplementedContentCheckServiceServer
}

func (s *GRPCServer) CheckHealth(ctx context.Context, request *api.EmptyRequest) (*api.HealthResponse, error) {
	fmt.Print("CHECKED")
	return &api.HealthResponse{}, nil
}
func RunRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}

}
func RunGrpc() {
	s := grpc.NewServer()
	srv := &GRPCServer{}
	api2.RegisterContentCheckServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
