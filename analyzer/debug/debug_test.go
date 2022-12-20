package debug

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEnable(t *testing.T) {
	defer func() {
		os.Setenv("DEBUG", "")
		logrus.SetLevel(logrus.InfoLevel)
	}()
	Enable()
	if os.Getenv("DEBUG") != "1" {
		t.Fatalf("expected DEBUG=1, got %s\n", os.Getenv("DEBUG"))
	}
	if logrus.GetLevel() != logrus.DebugLevel {
		t.Fatalf("expected log level %v, got %v\n", logrus.DebugLevel, logrus.GetLevel())
	}
}
