package db

import (
	"fmt"
	"log"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/config"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB Instance
var DB *gorm.DB

func New() {
	CreateDB()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Config.Postgres.Host, config.Config.Postgres.Port, config.Config.Postgres.User, config.Config.Postgres.Pass, config.Config.Postgres.Database)
	fmt.Println(dsn)
	var err error
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(fmt.Errorf("fatal error : DB Connection could not be established. %v", err))
	}

	Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Panic("Database migration COULD NOT be completed successfully.")
	}
	log.Println("Database migration completed successfully.")
}

func CreateDB() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Config.Postgres.Host, config.Config.Postgres.Port, config.Config.Postgres.User, config.Config.Postgres.Pass, "postgres")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	res := db.Exec("CREATE DATABASE " + config.Config.Postgres.Database + ";")
	if res.Error != nil {
		log.Printf("database %s already exists", config.Config.Postgres.Database)
	}
}
