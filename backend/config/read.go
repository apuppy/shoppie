package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf Configurations

func init() {
	// Set the file name of the configurations file
	viper.SetConfigName("settings")

	// Set configuration file format
	viper.SetConfigType("yml")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config/")


	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	err := viper.Unmarshal(&Conf)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}
}
