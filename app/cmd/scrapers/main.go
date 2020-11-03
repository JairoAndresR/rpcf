package main

import (
	"github.com/joho/godotenv"
	"log"
	"rpcf/app/dataproviders/queue"
	"rpcf/core"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var q queue.Client
	invokeFunc := func(qq queue.Client) {
		q = qq
	}

	err = core.Injector.Invoke(invokeFunc)
	if err != nil {
		panic(err)
	}

	w := newProductsWorker(q)
	w.Execute()
}
