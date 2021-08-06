package helper

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(level string) *logrus.Logger {
	fmt.Println("loglevel:", level)
	logruslevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	log := &logrus.Logger{
		Out:          os.Stdout,
		Hooks:        make(logrus.LevelHooks),
		Level:        logruslevel,
		ReportCaller: false,
	}
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})
	return log
}
