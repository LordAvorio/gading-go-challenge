package main

import (
	"fmt"
	"gading-go-challenge/database"
	"gading-go-challenge/routes"
	"gading-go-challenge/utils"

	"github.com/spf13/viper"
)

func init(){
	utils.ReadConfigEnv()
}

func main() {

	db, err := database.StartDB()
	if err != nil {
		panic(err)
	}

	runMigration := viper.GetBool("AUTO_MIGRATE")

	if runMigration {
		database.RunMigration(db)
	}

	port := viper.GetString("APP_PORT")

	app:= routes.RouteSession(db)
	app.Run(port)

	fmt.Println("SUCCESS RUN")
}