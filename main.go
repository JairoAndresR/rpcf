package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"rpcf/app/delivery/rest"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := rest.NewServer()
	s.Run()
}
