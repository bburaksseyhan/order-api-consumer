package main

import (
	"github.com/bburaksseyhan/orderconsumer/src/cmd/utils"
	services "github.com/bburaksseyhan/orderconsumer/src/pkg"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {
	log.Info("main.go is running...")

	settings := read()

	services.Initialize(settings)
}

func read() utils.Configuration {
	//Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var config utils.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Error("Unable to decode into struct, %v", err)
	}

	return config
}
