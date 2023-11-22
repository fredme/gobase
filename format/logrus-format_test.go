package format

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

/*
output:
2023-11-22 15:20:55 [info] [logrus-format_test.go:16] info message
2023-11-22 15:20:55 [error] [logrus-format_test.go:17] error message
*/
func TestLogRusFormat(t *testing.T) {
	logger := &logrus.Logger{
		Out:       os.Stdout,
		Level:     logrus.DebugLevel,
		Formatter: &MyLogFormatter{},
	}
	logger.Infof("info message")
	logger.Errorf("error message")
}
