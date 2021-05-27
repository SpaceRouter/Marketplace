package main

import (
	"fmt"
	"github.com/spacerouter/marketplace/config"
	"github.com/spacerouter/marketplace/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
	conf := config.GetConfig()
	dbType := conf.GetString("db.type")

	if dbType == "sqlite" {
		db, err := gorm.Open(sqlite.Open(conf.GetString("db.database")), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	} else if dbType == "mysql" {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: conf.GetString("db.user") + ":" + conf.GetString("db.password") + "@tcp(" + conf.GetString("db.host") + ":" + conf.GetString("db.port") + ")/" + conf.GetString("db.database") + "?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		}), &gorm.Config{})
		return db, err
	}
	return nil, fmt.Errorf("db type not yet implemented %s", dbType)
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Developer{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Stack{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Service{}, &models.VolumeDeclaration{}, &models.NetworkDeclaration{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Developer{}, &models.Volume{}, &models.EnvVar{}, models.Network{}, models.Port{})
	if err != nil {
		return err
	}

	return nil
}
