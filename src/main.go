package main

import (
	"flag"
	"fmt"
	"github.com/spacerouter/marketplace/config"
	"github.com/spacerouter/marketplace/models"
	"github.com/spacerouter/marketplace/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Developer{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Stack{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Service{}, &models.VolumesDeclaration{}, &models.NetworkDeclaration{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Developer{}, &models.Volume{}, &models.EnvVar{}, models.Network{})
	if err != nil {
		log.Fatal(err)
	}

	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}

}
