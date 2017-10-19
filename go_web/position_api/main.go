package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=yunfeng_db dbname=position sslmode=disable password=mypassword")
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	defer db.Close()
}
