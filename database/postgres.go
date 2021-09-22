package database

import (
	"fmt"
	"log"

	"github.com/ditointernet/tradulab-service/drivers"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func MustNewDB() Database {
	return Database{}
}

func (d *Database) AutoMigration(arg ...interface{}) error {
	err := d.db.AutoMigrate(arg...)

	if err != nil {
		return errors.Wrap(err, "Couldnt migrate database structs")
	}

	return nil
}

func (d *Database) StartPostgres() *gorm.DB {
	// colocar em variavel de ambiente
	dns := "host=tradulab-db user=admin password=12345 dbname=tradulab port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("error: ", err)
	}

	d.db = database

	fmt.Println(d.db, "------------------")

	return nil

}

func (d *Database) GetDatabase() *gorm.DB {
	return d.db
}

// mudar para domains vai vir do database agr
func (d *Database) SaveFile(file *drivers.File) error {
	db := d.GetDatabase()

	fmt.Println(db, "-------------------")

	fmt.Println(file.ID)
	query := db.Exec(
		"INSERT into files (id, project_id, file_path) values (?,?,?)",
		file.ID,
		file.ProjectID,
		file.FilePath,
	)

	return query.Error
}
