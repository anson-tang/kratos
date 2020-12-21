package log

import (
	"testing"
)

func TestLogger(t *testing.T) {
	log := GetLogger("test")
	log.Print("log", "test")

	Debug(log).Print("log", "test")
	Info(log).Print("log", "test")
	Warn(log).Print("log", "test")
	Error(log).Print("log", "test")
}