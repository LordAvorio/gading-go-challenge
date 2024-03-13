package utils

import "github.com/spf13/viper"

func ReadConfigEnv() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if errReadConfig := viper.ReadInConfig(); errReadConfig != nil {
		panic(errReadConfig)
	}

}