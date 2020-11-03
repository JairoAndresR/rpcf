package main

import (
	"rpcf/app/dataproviders/queue"
	"rpcf/core"
)

func main() {

	var queue queue.Client
	invokeFunc := func(queue queue.Client) {
		queue = queue
	}

	err := core.Injector.Invoke(invokeFunc)
	if err != nil {
		panic(err)
	}

	w := newProductsWorker(queue)
	w.Execute()
}
