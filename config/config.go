package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	User          string
	Password      string
	Host          string
	Port          int
	DbName        string
	EnableSslMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DBConfig     *DbConfig
}

var configurations *Config

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

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT secret key is required")
		os.Exit(1)
	}

	db_user := os.Getenv("DB_USER")
	if db_user == "" {
		fmt.Println("username is required")
		os.Exit(1)
	}

	db_password := os.Getenv("DB_PASSWORD")
	if db_password == "" {
		fmt.Println("password is required")
		os.Exit(1)
	}

	db_host := os.Getenv("DB_HOST")
	if db_host == "" {
		fmt.Println("Host is required")
		os.Exit(1)
	}

	db_port_as_str := os.Getenv("DB_PORT")
	if db_port_as_str == "" {
		fmt.Println("Database port is required")
		os.Exit(1)
	}

	db_port, err := strconv.Atoi(db_port_as_str)
	if err != nil {
		fmt.Println("Port number must be integer")
		os.Exit(1)
	}

	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		fmt.Println("Database name is required")
		os.Exit(1)
	}

	enbl_ssl_mode := os.Getenv("ENABLE_SSL_MODE")

	enable_ssl_mode, err := strconv.ParseBool(enbl_ssl_mode)
	if err != nil {
		fmt.Println("SSL mode value should be boolean")
		os.Exit(1)
	}

	dbCnf := &DbConfig{
		User:          db_user,
		Password:      db_password,
		Host:          db_host,
		Port:          db_port,
		DbName:        db_name,
		EnableSslMode: enable_ssl_mode,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     httpPort,
		JwtSecretKey: jwtSecretKey,
		DBConfig:     dbCnf,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
