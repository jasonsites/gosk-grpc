package resolver

import (
	"fmt"

	"github.com/jasonsites/gosk-grpc/config"
	"github.com/jasonsites/gosk-grpc/internal/grpcserver"
	"github.com/sirupsen/logrus"
)

// Config provides a singleton config.Configuration instance
func (r *Resolver) Config() *config.Configuration {
	if r.config == nil {
		c, err := config.LoadConfiguration()
		if err != nil {
			panic(fmt.Errorf("error resolving config: %v", err))
		}

		r.config = c
	}

	return r.config
}

// GrpcServer provides a singleton grpcserver.Server instance
func (r *Resolver) GrpcServer() *grpcserver.Server {
	if r.grpcServer == nil {
		server, err := grpcserver.NewServer(&grpcserver.Config{
			HealthServer: r.HealthServer(),
			Host:         r.Config().Server.Host,
			Log:          r.Log(),
			Port:         r.Config().Server.Port,
		})
		if err != nil {
			panic(fmt.Errorf("error resolving grpc server: %v", err))
		}

		r.grpcServer = server
	}

	return r.grpcServer
}

// HealthServer provides a singleton grpcserver.HealthServer instance
func (r *Resolver) HealthServer() *grpcserver.HealthServer {
	if r.healthServer == nil {
		server, err := grpcserver.NewHealthServer(&grpcserver.HSConfig{
			Log: r.Log(),
		})
		if err != nil {
			panic(fmt.Errorf("error resolving grpc health server: %v", err))
		}

		r.healthServer = server
	}

	return r.healthServer
}

// Log provides a singleton logrus.FieldLogger instance
func (r *Resolver) Log() logrus.FieldLogger {
	if r.log == nil {
		r.log = logrus.WithFields(logrus.Fields{
			"name":    r.metadata.Name,
			"version": r.metadata.Version,
		})

		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyLevel: "plevel",
			},
		})

		level, err := logrus.ParseLevel(r.Config().Log.Level)
		if err != nil {
			level = logrus.InfoLevel
		}
		logrus.SetLevel(level)
	}

	return r.log
}
