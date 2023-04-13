package database

import (
	"fmt"
	"log"
	"tugas-sesi12/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = "5432"
	username = "postgres"
	password = "root123"
	dbname   = "h8-products"
	ssl      = "disable"
	dialect  = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func handleDatabaseConnection() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", host, username, password, dbname, port, ssl)
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

	alterConstraint :=
		`
	ALTER TABLE "public"."products" 
		ADD CONSTRAINT "products_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
	`
	db.Exec(alterConstraint)

}

func InitializeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTable()
}

func GetDatabaseInstance() *gorm.DB {
	return db
}
