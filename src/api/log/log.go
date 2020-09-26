package log

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/config"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	Log = &logrus.Logger{
		Out:   os.Stdout,
		Level: level,
	}

	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
	}
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	// fields are extra fields added to the log output
	Log.WithFields(parseFields(tags...)).Info(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, field := range tags {
		splits := strings.Split(field, ":")
		result[strings.TrimSpace(splits[0])] = strings.TrimSpace(splits[1])
	}
	return result
}
