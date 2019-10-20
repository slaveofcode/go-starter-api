package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func setup(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.WarnLevel)
}

// New logger from logrus
func New() *logrus.Logger {
	logger := logrus.New()
	setup(logger)
	return logger
}
