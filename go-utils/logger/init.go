package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetupLogger(logLevel string) (*logrus.Logger, error) {
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	Log := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     lvl,
	}
	return Log, nil
}
