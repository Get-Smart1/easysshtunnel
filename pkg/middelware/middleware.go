package middelware

import (
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

func (m *Middleware) CreateNewConnection(info connection.ConnectionInfo) {
	if info.Provider == "" {

	}
}

func (m *Middleware) UpdateConnection(info connection.ConnectionInfo) {

}

func (m *Middleware) RemoveConnection(info connection.ConnectionInfo) {

}
