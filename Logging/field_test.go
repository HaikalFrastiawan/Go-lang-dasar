package golang

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("name","Haikal").Info("Hello World")

	logger.WithField("username", "frastiawan").WithField("name", "fras").Info("User logged in")


}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username" : "ekal",
		"name" : "kal",
	}).Info("ini pesan")
}