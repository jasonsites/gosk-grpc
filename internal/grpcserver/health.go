package grpcserver

import (
	"context"

	"github.com/jasonsites/gosk-grpc/internal/validation"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// HSConfig defines the input to NewHealthServer
type HSConfig struct {
	Log logrus.FieldLogger `validate:"required"`
}

type HealthServer struct {
	log logrus.FieldLogger `validate:"required"`
}

func NewHealthServer(c *HSConfig) (*HealthServer, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	return &HealthServer{log: c.Log}, nil
}

func (h *HealthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	h.log.Infof("server health check")

	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthServer) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	h.log.Infof("server health watch")

	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}
