package main

import (
	"fmt"

	"github.com/ditointernet/tradulab-service/repository"
)

func main() {

	env, pan := repository.GoDotEnvVariable()
	if pan != nil {
		fmt.Println("Error during environment variables build", pan.Error())
		return
	}

	db := repository.NewConfig(&repository.ConfigDB{
		User:     env.User,
		Host:     env.Host,
		Password: env.Password,
		DbName:   env.DbName,
		Port:     env.Port,
	})

	tables := &repository.File{}

	db.StartPostgres()

	err := db.AutoMigration(tables)

	if err != nil {
		panic(err)
	}

}
