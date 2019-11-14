package core

import (
	"os"

	"github.com/op/go-logging"
)

var (
	logLevel = GetEnvOrDefault("LogLevel", "INFO")
)

// InitLogger initilaze the logging driver. Also setted log level from
// 'LogLevel' environment variable.
func InitLogger() {
	// TODO: Panic yakala

	_ = logging.MustGetLogger("base")
	stdout := logging.NewLogBackend(os.Stdout, "", 0)

	format := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.5s} %{id:03x}%{color:reset} %{message}`)
	logging.SetFormatter(format)

	levelBackend := logging.AddModuleLevel(stdout)

	switch logLevel {
	case "CRITICAL":
		levelBackend.SetLevel(logging.CRITICAL, "")
	case "ERROR":
		levelBackend.SetLevel(logging.ERROR, "")
	case "WARN":
		levelBackend.SetLevel(logging.WARNING, "")
	case "NOTICE":
		levelBackend.SetLevel(logging.NOTICE, "")
	case "DEBUG":
		levelBackend.SetLevel(logging.DEBUG, "")
	default:
		levelBackend.SetLevel(logging.INFO, "")
	}

	logging.SetBackend(levelBackend)

}
