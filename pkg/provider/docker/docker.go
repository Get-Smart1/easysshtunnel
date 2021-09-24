package docker

import (
	"context"
	"easytunnel/pkg/connection"
	log2 "easytunnel/pkg/log"
	"easytunnel/pkg/middelware"
	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/docker/docker/client"
	"strings"
)

const (
	UpdateInterval            = 3
	EasyTunnelMiddlewareLabel = "easytunnel.middleware"
)

var (
	log logrus.Logger = *log2.GetLogger()
)

type container struct {
	id     string
	labels map[string]string
	ports  []int
}

type containerList []container

type Docker struct {
	ctx context.Context
	cli *client.Client
}

func (docker *Docker) Initialize() {
	log.Info("Initialize docker provider")

	docker.ctx = context.Background()
	var err error
	docker.cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(), client.WithHost("unix:///var/run/docker.sock"))

	if err != nil {
		log.Error(err)
		panic(err)
	}

	var containers containerList

	for {
		containers.addNewContainers(docker.getAllContainers())
		containers.updateContainers(docker.getAllContainers())
		containers.removeContainers(docker.getAllContainers())
		time.Sleep(UpdateInterval * time.Second)
	}

}

//Get all Containers with easytunnel label
func (docker *Docker) getAllContainers() containerList {

	var result containerList

	containers, err := docker.cli.ContainerList(docker.ctx, types.ContainerListOptions{})
	if err != nil {
		log.Error(err)
		panic(err)
	}

	for _, item := range containers {
		labels := getEasyTunnelLabelsFromContainer(item.Labels)
		if len(labels) == 0 {
			continue
		}

		var container container
		container.labels = labels
		container.id = item.ID
		result = append(result, container)
	}

	return result
}

func getEasyTunnelLabelsFromContainer(labels map[string]string) map[string]string {
	result := make(map[string]string)

	for key, element := range labels {
		if strings.HasPrefix(key, "easytunnel") {
			result[key] = element
		}
	}

	return result
}

func (list *containerList) addNewContainers(containers containerList) {

	for _, item := range containers {
		if !list.containsID(item.id) {
			*list = append(*list, item)
			middelware.CreateNewConnection(item.getMiddleware(), item.getConnectionInfo())
		}
	}
}

func (list *containerList) updateContainers(containers containerList) {
	for i, cItem := range *list {
		for _, nItem := range containers {
			if cItem.id == nItem.id {
				if !cItem.equals(nItem) {
					middelware.UpdateConnection(nItem.getMiddleware(), nItem.getConnectionInfo())
					*list = remove(*list, i)
				}
			}
		}
	}
}

func (list *containerList) removeContainers(containers containerList) {

}

func (list containerList) containsID(id string) bool {

	for _, item := range list {
		if item.id == id {
			return true
		}
	}
	return false
}

func (container container) getMiddleware() string {

	for key, value := range container.labels {
		if key == EasyTunnelMiddlewareLabel {
			return value
		}
	}
	return ""
}

func (container container) getConnectionInfo() connection.ConnectionInfo {
	var result connection.ConnectionInfo

	return result
}

func remove(list containerList, index int) containerList {
	list[index] = list[len(list)-1]
	return list[:len(list)-1]
}

func (c container) equals(other container) bool {
	if c.id != other.id {
		return false
	}

	if len(c.ports) != len(other.ports) {
		return false
	}
	for _, port := range c.ports {
		found := false
		for _, otherport := range other.ports {
			if port == otherport {
				found = true
			}
		}
		if !found {
			return false
		}
	}

	if len(c.labels) != len(other.labels) {
		return false
	}

	for cKey, cElement := range c.labels {
		different := true
		for nKey, nElement := range other.labels {
			if cKey == nKey {
				if cElement == nElement {
					different = false
				}
			}
		}
		if different {
			return false
		}
	}

	return true
}
