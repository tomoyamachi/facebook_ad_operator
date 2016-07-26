package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("app")    // no need to include file extension
	viper.AddConfigPath("config") // set the path of your config file

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		devServer := viper.GetString("development.server")
		devConnectionMax := viper.GetInt("development.connection_max")
		devEnabled := viper.GetBool("development.enabled")
		devPort := viper.GetInt("development.port")

		fmt.Printf("\nDevelopment Config found:\n server = %s\n connection_max = %d\n"+
			" enabled = %t\n"+
			" port = %d\n",
			devServer,
			devConnectionMax,
			devEnabled,
			devPort)

		prodServer := viper.GetString("production.server")
		prodConnectionMax := viper.GetInt("production.connection_max")
		prodEnabled := viper.GetBool("production.enabled")
		prodPort := viper.GetInt("production.port")

		fmt.Printf("\nProduction Config found:\n server = %s\n connection_max = %d\n"+
			" enabled = %t\n"+
			" port = %d\n",
			prodServer,
			prodConnectionMax,
			prodEnabled,
			prodPort)
	}

}
