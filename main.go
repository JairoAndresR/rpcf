package main

import (
	"rpcf/app/delivery/rest"
)

func main() {
	s := rest.NewServer()
	s.Run()
}
