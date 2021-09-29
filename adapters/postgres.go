package adapters

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
	in *Config
}

func NewDatabase(in *Config) *Database {
	return &Database{
		in: in,
	}
}

func (d *Database) AutoMigration(arg ...interface{}) error {
	err := d.db.AutoMigrate(arg...)

	if err != nil {
		return errors.Wrap(err, "Couldnt migrate database structs")
	}

	return nil
}

func (d *Database) StartPostgres() {

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		d.in.Host,
		d.in.User,
		d.in.Password,
		d.in.DbName,
		d.in.Port,
	)

	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("error: ", err)
	}

	d.db = database
}

func (d *Database) GetDatabase() *gorm.DB {
	return d.db
}
