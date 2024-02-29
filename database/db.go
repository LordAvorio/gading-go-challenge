package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	db  *sql.DB
	err error
)

func StartDB() {

	serverHost := viper.GetString("HOST")
	serverPort := viper.GetString("PORT")
	databaseUsername := viper.GetString("DB_USER")
	databasePassword := viper.GetString("DB_PASS")
	databaseName := viper.GetString("DB_NAME")
	databaseType := viper.GetString("DB_TYPE")

	mySqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8", databaseUsername, databasePassword, serverHost, serverPort, databaseName)

	db, err = sql.Open(databaseType, mySqlInfo)
	if err != nil {
		panic(err)
	}
	

	fmt.Println("Successfully connected to database")

}

func GetDB() *sql.DB {
	return db
}
