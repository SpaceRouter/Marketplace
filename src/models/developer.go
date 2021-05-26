package models

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	Name    string
	Website string
}
