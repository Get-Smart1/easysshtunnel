package main

import (
	"easytunnel/pkg/middelware"
	"easytunnel/pkg/middelware/sshdocker"
	"easytunnel/pkg/provider/docker"
	"time"
)

func main() {

	var middelware sshdocker.SshDocker
	var provider docker.Docker

	go provider.Initialize(&middelware)

	time.Sleep(500 * time.Second)

}

func Lol(middleware *middelware.IMiddleware) {

}
