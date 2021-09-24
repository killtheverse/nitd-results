package config

import (
	"os"
)

// Server configeration structure
type Config struct {
	ServerAddress	string
	DBUser			string
	DBPassword		string
	DBName			string
	DBURI			string	// mongodb connection URI
}

// initialize will read the environment variables and store them in the config struct
func (config *Config) initialize() {
	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	config.DBUser = os.Getenv("MONGODB_USER")
	config.DBPassword = os.Getenv("MONGODB_PASSWORD")
	config.DBName = os.Getenv("MONGODB_NAME")
	config.DBURI = os.Getenv("MONGODB_URI")
}

// NewConfig will initialize and return the config
func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}