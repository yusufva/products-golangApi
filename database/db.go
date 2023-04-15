package database

import (
	"fmt"
	"log"
	"os"
	"tugas-sesi12/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dialect = "postgres"

	// uncomment this for ruby on rails
	// host     = os.Getenv("PG_HOST")
	// port     = os.Getenv("PG_PORT")
	// username = os.Getenv("PG_USERNAME")
	// password = os.Getenv("PG_PASSWORD")
	// dbname   = os.Getenv("PG_DBNAME")
	// ssl      = os.Getenv("PG_SSLMODE")

	//uncomment variable bellow for local using .env file
	host     = goDotEnvVariable("PG_HOST")
	port     = goDotEnvVariable("PG_PORT")
	username = goDotEnvVariable("PG_USERNAME")
	password = goDotEnvVariable("PG_PASSWORD")
	dbname   = goDotEnvVariable("PG_DBNAME")
	ssl      = goDotEnvVariable("PG_SSLMODE")
)

var (
	db  *gorm.DB
	err error
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Panicf("error while getting .env file : %s", err.Error())
	}

	return os.Getenv(key)
}

func handleDatabaseConnection() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta", host, port, username, password, dbname, ssl)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("error while connecting to db: %s", err.Error())
	}
}

func handleCreateRequiredTable() {
	err = db.Debug().AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		log.Panicln(err.Error())
	}

	// alterConstraint :=
	// 	`
	// ALTER TABLE "public"."products"
	// 	ADD CONSTRAINT "products_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
	// `
	// db.Exec(alterConstraint)

}

func InitializeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTable()
}

func GetDatabaseInstance() *gorm.DB {
	return db
}
