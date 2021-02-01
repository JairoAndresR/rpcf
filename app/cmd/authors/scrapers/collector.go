package main

import (
	_ "rpcf/app/products/adapters/repositories"
	"rpcf/core"
	_ "rpcf/products/managers"
	"rpcf/products/ports"
)

func NewCollector() ports.AuthorsCollector {
	var collector ports.AuthorsCollector
	invokeFunc := func(c ports.AuthorsCollector) {
		collector = c
	}

	err := core.Injector.Invoke(invokeFunc)
	if err != nil {
		panic(err)
	}

	return collector
}
