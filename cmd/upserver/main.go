package main

import (
	"upserver/internal/apiserver"
	log "github.com/sirupsen/logrus"
	"runtime"
)

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	log.WithFields(log.Fields{"cpuNum": cores}).Info("CPUNUM:")
	go apiserver.HttpStart()
	select {}
}
