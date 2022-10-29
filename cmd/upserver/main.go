package main

import (
	"upserver/internal/apiserver"
	"upserver/internal/pkg/service"
	"upserver/internal/pkg/utils"
)

func main() {
	utils.InitConfig()
	service.Init()
	apiserver.HttpStart()
}
