package main

import (
	"flag"
	"fmt"
	"github.com/spacerouter/marketplace/config"
	"github.com/spacerouter/marketplace/server"
	"log"
	"os"
)

// @title SpaceRouter Marketplace
// @version 0.1
// @description

// @contact.name ESIEESPACE Network
// @contact.url http://esieespace.fr
// @contact.email contact@esieespace.fr

// @license.name GPL-3.0
// @license.url https://github.com/SpaceRouter/authentication_server/blob/louis/LICENSE

// @host localhost:8081
// @BasePath /
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode} [COMMAND]")
		fmt.Println("Command list (default: Server):")
		fmt.Println("\tServer : launch server")
		fmt.Println("\tImport : import docker compose")
		fmt.Println("\tCreateDev : create new developer")

		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	switch flag.Arg(0) {
	case "Import":

		ImportCompose(flag.Arg(1), db)
		return
	case "Server", "":
		err = server.Init(db)
		if err != nil {
			log.Fatal(err)
		}
		return
	case "CreateDev":
		createDeveloper(db)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
