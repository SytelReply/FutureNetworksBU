package common

import (
	"os"
)

type ConfigStore struct {
	config *Config
}

type Config struct {
	GrpcPort    string
	GrpcAddress string
	RestPort    string
}

func GetConfig() *Config {

	GRPC_SERVER_PORT := os.Getenv("GRPC_SERVER_PORT")
	GRCP_SERVER_NETWORK_ADDRESS := os.Getenv("GRPC_NETWORK_ADDRESS")
	REST_SERVER_PORT := os.Getenv("REST_SERVER_PORT")

	if GRPC_SERVER_PORT == "" {
		GRPC_SERVER_PORT = "8080"
	}
	if GRCP_SERVER_NETWORK_ADDRESS == "" {
		GRCP_SERVER_NETWORK_ADDRESS = "localhost"
	}
	if REST_SERVER_PORT == "" {
		REST_SERVER_PORT = "8081"
	}
	return &Config{
		GrpcPort:    GRPC_SERVER_PORT,
		GrpcAddress: GRCP_SERVER_NETWORK_ADDRESS,
		RestPort:    REST_SERVER_PORT,
	}
}
