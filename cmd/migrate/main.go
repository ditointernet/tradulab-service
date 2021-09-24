package main

import (
	"fmt"

	"github.com/ditointernet/tradulab-service/database"
)

func main() {

	env, pan := database.GoDotEnvVariable()
	if pan != nil {
		fmt.Println("Error during environment variables build", pan.Error())
		return
	}

	db := database.NewConfig(&database.ConfigDB{
		User:     env.User,
		Host:     env.Host,
		Password: env.Password,
		DbName:   env.DbName,
		Port:     env.Port,
	})

	tables := &database.File{}

	db.StartPostgres()

	err := db.AutoMigration(tables)

	if err != nil {
		panic(err)
	}

}
