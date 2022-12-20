package debug

import (
	"os"

	"github.com/sirupsen/logrus"
)

// enableDebug sets the DEBUG env var to true
// and makes the logger to log at debug level.
func Enable() {
	os.Setenv("DEBUG", "1")
	logrus.SetLevel(logrus.DebugLevel)
}

// disableDebug sets the DEBUG env var to false
// and makes the logger to log at info level.
func Disable() {
	os.Setenv("DEBUG", "")
	logrus.SetLevel(logrus.InfoLevel)
}

// isDebugEnabled checks whether the debug flag is set or not.
func IsEnabled() bool {
	return os.Getenv("DEBUG") != ""
}
