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

type Middleware struct {
	middlewares map[string]IMiddleware
}

func New() *Middleware {
	var result Middleware
	result.middlewares = make(map[string]IMiddleware)
	return &result
}

func (m *Middleware) AddMiddleware(name string, middleware IMiddleware) {
	m.middlewares[name] = middleware
}

func (m *Middleware) CreateNewConnection(middleware string, info connection.ConnectionInfo) {
	if middleware == "" {
		middleware = config.GetStringValue(config.DefaultMiddleware)
	}

}

func (m *Middleware) UpdateConnection(middleware string, info connection.ConnectionInfo) {

}

func (m *Middleware) RemoveConnection(middleware string, info connection.ConnectionInfo) {

}
