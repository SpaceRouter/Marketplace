package main

import (
	"github.com/spacerouter/marketplace/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
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
