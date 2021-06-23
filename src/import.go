package main

import (
	"bufio"
	"fmt"
	"github.com/spacerouter/marketplace/models"
	"github.com/spacerouter/marketplace/utils"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

func ImportCompose(path string, db *gorm.DB) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name:  ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Description:  ")
	desc, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Icon:  ")
	icon, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	compose := models.Compose{}

	err = yaml.Unmarshal(file, &compose)
	if err != nil {
		log.Fatal(err)
	}

	stack, err := utils.ConvertToStack(compose)
	if err != nil {
		log.Fatal(err)
	}

	stack.Name = name
	stack.Description = desc
	stack.Icon = icon

	db.Create(stack)
}

func RemoveStack(id string, db *gorm.DB) {
	db.Delete(&models.Stack{}, id)
}
