package middelware

import (
	"easytunnel/pkg/config"
	"easytunnel/pkg/connection"
)

type IMiddleware interface {
	Initialize()
	CreateNewConnection(connection connection.ConnectionInfo)
	UpdateConnection(connection connection.ConnectionInfo)
	RemoveConnection(connection connection.ConnectionInfo)
}

var (
	middlewares map[string]IMiddleware
)

func init() {
	middlewares = make(map[string]IMiddleware)
}

func AddMiddleware(name string, middleware IMiddleware) {
	middlewares[name] = middleware
}

func CreateNewConnection(middleware string, info connection.ConnectionInfo) {
	if middleware == "" {
		middleware = config.GetStringValue(config.DefaultMiddleware)
	}
	middlewares[middleware].CreateNewConnection(info)
}

func UpdateConnection(middleware string, info connection.ConnectionInfo) {

}

func RemoveConnection(middleware string, info connection.ConnectionInfo) {

}
