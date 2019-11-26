package core

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	level := []string{"CRITICAL", "ERROR", "WARN", "NOTICE", "DEBUG", "NONE"}
	for i := 0; i < len(level); i++ {
		logLevel = level[i]
		go InitLogger()
	}
}
