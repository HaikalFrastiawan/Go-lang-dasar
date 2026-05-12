package golang

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("This is a hook", entry.Level, entry.Message)
	return nil
}


func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("Hello logging")
	logger.Warn("Hello logging")
	logger.Error("Hello logging")
}