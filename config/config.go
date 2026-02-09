package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var configurations Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables ", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	portAsString := os.Getenv("HTTP_PORT")
	if portAsString == "" {
		fmt.Println("Port is required")
		os.Exit(1)
	}

	httpPort, err := strconv.Atoi(portAsString)
	if err != nil {
		fmt.Println("Port number must be integer")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}
