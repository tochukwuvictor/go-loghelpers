package main

import (
	"errors"

	log "github.com/hashicorp/go-hclog"
)

func main() {}

// SetupLogger returns a logger initialized with the log level and program name
func SetupLogger(logLevel, programName string) (logger log.Logger, err error) {
	level := log.LevelFromString(logLevel)
	if level == 0 {
		return logger, errors.New("Invalid log level specified: " + logLevel)
	}

	return log.New(&log.LoggerOptions{
		Name:  programName,
		Level: level,
	}), nil
}
