package resolver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jasonsites/gosk-grpc/config"
	"github.com/jasonsites/gosk-grpc/internal/grpcserver"
	"github.com/sirupsen/logrus"
)

// Config defines the input to NewResolver
type Config struct {
	Config       *config.Configuration
	GrpcServer   *grpcserver.Server
	HealthServer *grpcserver.HealthServer
	Log          logrus.FieldLogger
	Metadata     *Metadata
}

// Application metadata
type Metadata struct {
	Name    string
	Version string
}

// Resolver provides singleton instances of app components
type Resolver struct {
	config       *config.Configuration
	grpcServer   *grpcserver.Server
	healthServer *grpcserver.HealthServer
	log          logrus.FieldLogger
	metadata     *Metadata
}

// NewResolver returns a new Resolver instance
func NewResolver(c *Config) *Resolver {
	if c == nil {
		c = &Config{}
	}

	r := &Resolver{
		config:       c.Config,
		grpcServer:   c.GrpcServer,
		healthServer: c.HealthServer,
		log:          c.Log,
		metadata:     c.Metadata,
	}

	r.Metadata()
	r.Config()
	r.Log()

	return r
}

// Metadata provides a singleton application Metadata instance
func (r *Resolver) Metadata() *Metadata {
	if r.metadata == nil {
		var metadata *Metadata

		jsondata, err := ioutil.ReadFile("/app/package.json")
		if err != nil {
			fmt.Printf("error reading package.json file, %v:", err)
		}

		if err := json.Unmarshal(jsondata, &metadata); err != nil {
			fmt.Printf("error unmarshalling package.json, %v:", err)
		}

		r.metadata = metadata
	}

	return r.metadata
}
