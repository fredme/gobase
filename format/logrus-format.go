package format

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type MyLogFormatter struct{}

func (f *MyLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取文件路径和所在行
	_, filename, line, _ := runtime.Caller(7)
	scriptName := filepath.Base(filename)

	message := fmt.Sprintf("%s [%s] [%s:%d] %s\n",
		entry.Time.Format("2006-01-02 15:04:05"),
		entry.Level.String(),
		scriptName,
		line,
		entry.Message)

	return []byte(message), nil
}
