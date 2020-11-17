package main

import (
	_ "rpcf/app/products/adapters/repositories"
	"rpcf/core"
	_ "rpcf/products/managers"
	"rpcf/products/ports"
)

func NewCollector() ports.ProductCollector {
	var collector ports.ProductCollector
	invokeFunc := func(c ports.ProductCollector) {
		collector = c
	}

	err := core.Injector.Invoke(invokeFunc)
	if err != nil {
		panic(err)
	}

	return collector
}
