package models

import "gorm.io/gorm"

type Stack struct {
	gorm.Model
	Name        string
	DeveloperID uint
	Services    []Service
	Networks    []NetworkDeclaration
	Volumes     []VolumesDeclaration
}

type Service struct {
	gorm.Model
	Image    string
	Port     []Port
	Envs     []EnvVar
	Volumes  []Volume
	Networks []Network
	StackID  uint
}

type Port struct {
	gorm.Model
	Port      int
	ServiceID uint
}

type VolumesDeclaration struct {
	gorm.Model
	Name    string
	StackID uint
}

type Volume struct {
	gorm.Model
	Name       string
	MountPoint string
	ServiceID  uint
}

type EnvVar struct {
	gorm.Model
	Name         string
	DefaultValue string
	ServiceID    uint
}

type NetworkDeclaration struct {
	gorm.Model
	Name    string
	StackID uint
}

type Network struct {
	gorm.Model
	Name      string
	ServiceID uint
}
