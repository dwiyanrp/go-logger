package main

import (
	"fmt"

	"github.com/dwiyanrp/logger"

	"github.com/spf13/viper"
)

func main() {
	// Setup Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	// Initialize our logger with Viper
	logger.Init(viper.GetViper())

	// Example usage
	logger.Debug("database", "This is a database-related debug message") // not shown
	logger.Debug("network", "This is a network-related debug message")   // not shown
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
