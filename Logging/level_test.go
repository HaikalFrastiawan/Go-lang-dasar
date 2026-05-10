package golang

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)



	logger.Trace("info")
	logger.Debug("info")
	logger.Info("info")
	logger.Warn("info")
	logger.Error("info")

}