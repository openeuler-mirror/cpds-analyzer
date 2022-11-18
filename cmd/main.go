package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initLogging() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	logrus.SetOutput(os.Stdout)
}

func main() {
	initLogging()

	logrus.Infof("Starting cpds-analyzer...")
}
