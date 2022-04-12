package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/jasonsites/gosk-grpc/internal/protos"
	"github.com/jasonsites/gosk-grpc/internal/utils"
	"github.com/jasonsites/gosk-grpc/internal/validation"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Config defines the input to NewServer
type Config struct {
	HealthServer *HealthServer `validate:"required"`
	Host         string
	Log          logrus.FieldLogger `validate:"required"`
	Port         uint               `validate:"required"`
}

// Server defines a gRPC server for handling domain action requests
type Server struct {
	protos.UnimplementedDomainServer
	healthServer *HealthServer
	host         string
	log          logrus.FieldLogger
	port         uint
	rpcServer    *grpc.Server
}

// NewServer returns a new Server instance
func NewServer(c *Config) (*Server, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	server := &Server{
		healthServer: c.HealthServer,
		host:         c.Host,
		log:          c.Log,
		port:         c.Port,
	}

	rpcServer := grpc.NewServer()
	protos.RegisterDomainServer(rpcServer, server)
	grpc_health_v1.RegisterHealthServer(rpcServer, server.healthServer)
	server.rpcServer = rpcServer

	return server, nil
}

// Action returns an DomainResponse protobuf message
func (s *Server) Action(ctx context.Context, req *protos.DomainRequest) (*protos.DomainResponse, error) {

	fmt.Printf("Request: %v\n", req)
	id := req.Prop.Id

	// ...path to application logic...

	response := &protos.DomainResponse{
		Status: id,
	}

	fmt.Printf("Response: %v\n", utils.ToJSON(response))

	return response, nil
}

// Listen starts a gRPC server on the configured port
func (s *Server) Listen() {
	s.log.Infof("grpc server listening on port %d", s.port)

	address := fmt.Sprintf("%s:%d", s.host, s.port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		s.log.Fatalf("failed to listen: %v", err)
	}

	s.Serve(lis)
}

// Serve starts a gRPC server on the configured port
func (s *Server) Serve(listener net.Listener) {
	if err := s.rpcServer.Serve(listener); err != nil {
		s.log.Fatalf("failed to serve: %v", err)
	}
}
