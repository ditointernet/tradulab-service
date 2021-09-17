package main

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectWithDataBase() {
	str := "host=localhost port=8000 user=postgres dbname=tradulab sslmode=disable password=abacate"

	fmt.Println(str)
}
