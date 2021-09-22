package sshdocker

import (
	"easytunnel/pkg/connection"
	log2 "easytunnel/pkg/log"
	"github.com/sirupsen/logrus"
)

type SshDocker struct {
	name string
}

var (
	log logrus.Logger = *log2.GetLogger()
)

func (*SshDocker) Initialize() {
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
