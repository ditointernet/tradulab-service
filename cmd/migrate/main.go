package main

import (
	"fmt"

	"github.com/ditointernet/tradulab-service/adapters"
	"github.com/ditointernet/tradulab-service/driven"
)

func main() {

	env, pan := adapters.GoDotEnvVariable()
	if pan != nil {
		fmt.Println("Error during environment variables build", pan.Error())
		return
	}

	db := adapters.NewDatabase(&adapters.Config{
		User:     env.User,
		Host:     env.Host,
		Password: env.Password,
		DbName:   env.DbName,
		Port:     env.Port,
	})

	tables := &driven.File{}

	db.StartPostgres()

	err := db.AutoMigration(tables)

	if err != nil {
		panic(err)
	}

}
