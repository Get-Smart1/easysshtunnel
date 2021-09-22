package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	mainLogger logrus.Logger
)

func init() {
	mainLogger = *logrus.StandardLogger()
	mainLogger.SetOutput(os.Stdout)
	SetLevel(logrus.DebugLevel)
}

func SetLevel(level logrus.Level) {
	mainLogger.SetLevel(level)
}

func GetLogger() *logrus.Logger {
	return &mainLogger
}
