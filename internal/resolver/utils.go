package resolver

import (
	"fmt"

	"github.com/jasonsites/gosk-grpc/config"
)

// postgresDSN returns a data source name string from a given postgres configuration
func postgresDSN(p config.Postgres) string {
	// TODO: sslmode
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.Database,
	)
}
