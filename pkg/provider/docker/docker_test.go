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
ad	update callBackUpdate
}

func (DummyMiddleware) Initialize() {
	panic("implement me")
}

func (middlware *DummyMiddleware) CreateNewConnection(connection connection.ConnectionInfo) {
	middlware.createNew()
}

func (middlware DummyMiddleware) UpdateConnection(connection connection.ConnectionInfo) {
	middlware.update()
}

func (DummyMiddleware) RemoveConnection(connection connection.ConnectionInfo) {
	panic("implement me")
}

func init_ContainersTest() (cList containerList, nList containerList, dummyMiddleware DummyMiddleware) {
	cList = make(containerList, 0)
	nList = make(containerList, 0)
	return cList, nList, dummyMiddleware
}

func Test_containerList_addNewContainers_AddedNoneContainer(t *testing.T) {

	cList, nList, dummyMiddleware := init_ContainersTest()
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

	cList, nList, dummyMiddleware := init_ContainersTest()

	callCounter := 0
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

func Test_containerList_updateContainers_ChangedNone(t *testing.T) {

	cList, nList, middleware := init_ContainersTest()

	middleware.update = func() {
		t.Error("no container was updated")
	}
	middelware.AddMiddleware("dummy", &middleware)

	cList = append(cList, container{
		id: "1",
		labels: map[string]string{
			"easytunnel.enable":       "true",
			EasyTunnelMiddlewareLabel: "dummy",
		},
	})

	nList = append(nList, container{
		id: "1",
		labels: map[string]string{
			"easytunnel.enable":       "true",
			EasyTunnelMiddlewareLabel: "dummy",
		},
	})

	cList.updateContainers(nList)

}


func Test_containerList_updateContainers_ChangedOneContainer(t *testing.T) {

	cList, nList, middleware := init_ContainersTest()

	counter := 0
	middleware.update = func() {
		counter++
	}
	middelware.AddMiddleware("dummy", &middleware)

	cList = append(cList, container{
		id: "1",
		labels: map[string]string{
			"easytunnel.enable":       "true",
			EasyTunnelMiddlewareLabel: "dummy",
		},
	})

	nList = append(nList, container{
		id: "1",
		labels: map[string]string{
			"easytunnel.enable":       "false",
			EasyTunnelMiddlewareLabel: "dummy",
		},
	})

	cList.updateContainers(nList)
	if counter == 0 {
		t.Error("Update on Middleware was not called")
	}
	assert.Equal(t, counter, 1, "The Update Container Function in middleware was called multiple times")



}