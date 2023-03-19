package log

import (
	"testing"
)

func init() {
	logger := Logger{
		Level:   "debug",
		Console: true,
		File:    nil,
	}

	Init(&logger)
}

func TestInfo(t *testing.T) {
	Info("test info.")
}

func TestDebug(t *testing.T) {
	Debug("test debug.")
}

func TestDebugf(t *testing.T) {
	Debugf("test debugf, parameter: %v", "value")
}

func TestWarnf(t *testing.T) {
	Warnf("test Warnf, parameter: %v", "warn value")
}
