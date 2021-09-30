package adapters

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(in *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		in.Host,
		in.User,
		in.Password,
		in.DbName,
		in.Port,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("error: ", err)
	}
	return database
}
