package database

import (
	"fmt"
	"log"

	"github.com/ditointernet/tradulab-service/drivers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func MustNewDB() Database {
	return Database{}
}

func (d *Database) StartPostgres() {
	dns := "host=database port=5432 user=admin dbname=tradulab sslmode=disable password=123456"

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

func (d *Database) SaveFile(file *drivers.File) error {
	db := d.GetDatabase()

	return db.Raw("INSERT into files (ID, ProjectID, FilePath) values (?,?,?)", file.ID, file.ProjectID, file.FilePath).Error

	//stmt.Exec(file.ID, file.ProjectID, file.ProjectID)

}
