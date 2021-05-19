package main

import (
	"flag"
	"fmt"
	"github.com/spacerouter/marketplace/config"
	"github.com/spacerouter/marketplace/server"
	"log"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Run as %s \n", user.Username)

	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)
	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}
}
