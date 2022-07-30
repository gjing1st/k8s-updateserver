package main

import (
	"upserver/internal/apiserver"
	"upserver/internal/pkg/service"
)

func main() {
	service.Init()
	go apiserver.HttpStart()

	select {}
}
