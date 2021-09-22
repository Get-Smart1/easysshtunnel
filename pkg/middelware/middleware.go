package middelware

import "easytunnel/pkg/connection"

type Middleware interface {
	Initialize()
	CreateNewConnection(connection connection.ConnectionInfo)
	UpdateConnection(connection connection.ConnectionInfo)
	RemoveConnection(connection connection.ConnectionInfo)
}
