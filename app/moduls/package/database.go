package moduls

import (
	"flag"
	"fmt"
	"bgin/app/libs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Conn *gorm.DB

func InitDB() *gorm.DB {
	dbhost := libs.GetEnvVariabel("DB_HOST", "127.0.0.1")
	dbport := libs.GetEnvVariabel("DB_PORT", "26257")
	dbuser := libs.GetEnvVariabel("DB_USER", "root")
	// dbpass := libs.GetEnvVariabel("DB_PASSWORD", "")
	dbname := libs.GetEnvVariabel("DB_NAME", "")
	var (
		configDB = flag.String("addr", fmt.Sprintf(
			"postgresql://%s@%s:%s/%s?sslmode=disable",
			dbuser, dbhost, dbport, dbname), "DB SETUP")
	)

	DB, err := gorm.Open("postgres", *configDB)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	return DB
}
