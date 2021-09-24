package docker

import (
	"easytunnel/pkg/connection"
	"easytunnel/pkg/middelware"
	"gotest.tools/assert"
	"testing"
)

type callbackCreateNew func()
type callBackUpdate func()

type DummyMiddleware struct {
	createNew callbackCreateNew
}

func (DummyMiddleware) Initialize() {
	panic("implement me")
}

func (middleware *DummyMiddleware) CreateNewConnection(connection connection.ConnectionInfo) {
	middleware.createNew()
}

func (DummyMiddleware) UpdateConnection(connection connection.ConnectionInfo) {
	panic("implement me")
}

func (DummyMiddleware) RemoveConnection(connection connection.ConnectionInfo) {
	panic("implement me")
}

func init_AddNewContainers() (cList containerList, nList containerList) {
	cList = make(containerList, 0)
	nList = make(containerList, 0)
	return cList, nList
}

func Test_containerList_addNewContainers_AddedNoneContainer(t *testing.T) {

	cList, nList := init_AddNewContainers()
	var dummyMiddleware DummyMiddleware
	dummyMiddleware.createNew = func() {
		t.Error("CreateNew was Called on Middleware, even when no new container was added")
	}
	middelware.AddMiddleware("dummy", &dummyMiddleware)
	cList.addNewContainers(nList)
	assert.Equal(t, len(cList), len(nList), "No item was added, length should be equal")

	nList = append(nList, container{
		id: "1",
		labels: map[string]string{
			"easytunnel.enable": "true",
		},
	})
}

func Test_containerList_addNewContainers_AddedOne(t *testing.T) {

	cList, nList := init_AddNewContainers()

	callCounter := 0
	var dummyMiddleware DummyMiddleware
	dummyMiddleware.createNew = func() {
		callCounter++
	}
	middelware.AddMiddleware("dummy", &dummyMiddleware)

	nList = append(nList, container{
		id: "1",
		labels: map[string]string{
			"easytunnel.enable":       "true",
			EasyTunnelMiddlewareLabel: "dummy",
		},
	})
	cList.addNewContainers(nList)

	if callCounter == 0 {
		t.Error("Middlware was not called")
	}
	assert.Equal(t, callCounter, 1, "CreateNewConnecton in Middleware was called multiple times")
	assert.Equal(t, len(cList), 1, "The Entry from new list, was not added to the current List")

}
