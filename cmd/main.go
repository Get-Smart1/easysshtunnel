package main

import (
	"easytunnel/pkg/middelware"
	"easytunnel/pkg/middelware/sshdocker"
)

func main() {

	sshDockerMiddleware := sshdocker.SshDocker{}
	sshDockerMiddleware.Initialize()
	middelware.AddMiddleware(&sshDockerMiddleware)

	for true {

	}

}

func Lol(middleware *middelware.IMiddleware) {

}
