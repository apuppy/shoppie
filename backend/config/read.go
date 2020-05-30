package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Read read configration from configration file
func Read() Settings {
	// setup config file read context
	SetConfContext()

	var Settings Settings
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	err := viper.Unmarshal(&Settings)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return Settings
}

// SetConfContext set configration file context before parsing, path
func SetConfContext() {
	// Set the file name of the configurations file
	viper.SetConfigName("settings")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config/")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Set configration file format
	viper.SetConfigType("yml")
}
