package database

import (
	"fmt"
	"log"

	"github.com/ditointernet/tradulab-service/drivers"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	User     string
	Host     string
	Password string
	DbName   string
	Port     string
}

type Database struct {
	db *gorm.DB
	in *ConfigDB
}

func MustNewDB() Database {
	return Database{}
}

func NewConfig(in *ConfigDB) *Database {
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
	// env, err := GoDotEnvVariable()
	// if err != nil {
	// 	fmt.Println("error when getting the environment variables")
	// 	return
	// }

	// host = ConfigDB.Host
	// user = ConfigDB.User
	// password = ConfigDB.Password
	// dbName = ConfigDB.DbName
	// port := ConfigDB.Port

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

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

// mudar para domains vai vir do database agr
func (d *Database) SaveFile(file *drivers.File) error {
	db := d.GetDatabase()

	query := db.Exec(
		"INSERT into files (id, project_id, file_path) values (?,?,?)",
		file.ID,
		file.ProjectID,
		file.FilePath,
	)

	return query.Error
}
