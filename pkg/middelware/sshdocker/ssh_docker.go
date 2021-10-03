package sshdocker

import (
	"bytes"
	"easytunnel/pkg/connection"
	log2 "easytunnel/pkg/log"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type SshDocker struct {
}

var (
	log         logrus.Logger = *log2.GetLogger()
	connections map[string]connection.ConnectionInfo
)

func (d *SshDocker) Initialize() {
	go d.UpdateConnectionStates()
}

func (d *SshDocker) CreateNewConnection(connection connection.ConnectionInfo) {
	log.Info("new connection")
}

func (d *SshDocker) UpdateConnection(connection connection.ConnectionInfo) {
	log.Info("new connection")
}

func (d *SshDocker) RemoveConnection(connection connection.ConnectionInfo) {
	log.Info("new connection")
}

func (d *SshDocker) GetName() string {
	return "ssh_docker"
}

func (d *SshDocker) UpdateConnectionStates() {
	jsonStr := []byte(`{"title":"Buy cheese and bread for breakfast."}`)

	req, err := http.NewRequest("POST", "http://139.162.142.90:8080", bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
