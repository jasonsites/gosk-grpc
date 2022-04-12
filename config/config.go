package config

import (
	"log"

	"github.com/spf13/viper"
)

// Configuration defines app configuration on startup
type Configuration struct {
	Log        Log      `validate:"required"`
	Postgres   Postgres `validate:"required"`
	Server     Server
	HttpServer Server
}

type Log struct {
	Level string `validate:"oneof=debug info warn error fatal"`
}

type Postgres struct {
	Database string `validate:"required"`
	Host     string `validate:"required"`
	Password string `validate:"required"`
	Port     uint   `validate:"max=65535"`
	User     string `validate:"required"`
}

type Server struct {
	Host string
	Port uint `validate:"max=65535"`
}

// LoadConfiguration loads config parameters on startup
func LoadConfiguration() (*Configuration, error) {
	var config Configuration

	viper.SetConfigName("config")

	viper.AddConfigPath("/app/config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.AllowEmptyEnv(true)

	// log
	if err := viper.BindEnv("log.level", "LOG_LEVEL"); err != nil {
		log.Fatalf("error binding env var `LOG_LEVEL`: %v", err)
	}

	// postgres
	if err := viper.BindEnv("postgres.database", "POSTGRES_DB"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_DB`: %v", err)
	}
	if err := viper.BindEnv("postgres.host", "POSTGRES_HOST"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_HOST`: %v", err)
	}
	if err := viper.BindEnv("postgres.password", "POSTGRES_PASSWORD"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_PASSWORD`: %v", err)
	}
	if err := viper.BindEnv("postgres.port", "POSTGRES_PORT"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_PORT`: %v", err)
	}
	if err := viper.BindEnv("postgres.user", "POSTGRES_USER"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_USER`: %v", err)
	}

	// server
	if err := viper.BindEnv("server.host", "SERVER_HOST"); err != nil {
		log.Fatalf("error binding env var `SERVER_HOST`: %v", err)
	}
	if err := viper.BindEnv("server.port", "SERVER_PORT"); err != nil {
		log.Fatalf("error binding env var `SERVER_PORT`: %v", err)
	}

	// http server
	if err := viper.BindEnv("httpserver.host", "HTTPSERVER_HOST"); err != nil {
		log.Fatalf("error binding env var `SERVER_HOST`: %v", err)
	}
	if err := viper.BindEnv("httpserver.port", "HTTPSERVER_PORT"); err != nil {
		log.Fatalf("error binding env var `SERVER_PORT`: %v", err)
	}

	// read and unmarshal config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error unmarshalling configuration: %v", err)
	}

	return &config, nil
}
