package main

import (
	"fmt"
	"github.com/spacerouter/marketplace/models"
	"gorm.io/gorm"
	"log"
)

func createDeveloper(db *gorm.DB) {
	fmt.Print("Name: ")
	name := ""
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Website: ")
	website := ""
	_, err = fmt.Scanln(&website)
	if err != nil {
		log.Fatal(err)
	}

	developer := models.Developer{
		Name:    name,
		Website: website,
	}

	result := db.Create(&developer)
	err = result.Error
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d rows affected\r\n", result.RowsAffected)
}
