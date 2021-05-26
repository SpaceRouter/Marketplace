package models

import "gorm.io/gorm"

type Stack struct {
	gorm.Model
	Name        string
	Icon        string
	Description string
	DeveloperID uint
	Services    []Service
	Networks    []NetworkDeclaration
	Volumes     []VolumeDeclaration
}

type Service struct {
	gorm.Model
	Name         string
	Image        string
	ImageVersion string
	Ports        []Port
	Envs         []EnvVar
	Volumes      []Volume
	Networks     []Network
	StackID      uint
}

type Port struct {
	gorm.Model
	Port      uint
	ServiceID uint
}

type VolumeDeclaration struct {
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
