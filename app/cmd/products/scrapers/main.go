package main

import (
	"github.com/joho/godotenv"
	"log"
	"rpcf/app/dataproviders/queue"
	"rpcf/core"
)

func main() {
	
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
