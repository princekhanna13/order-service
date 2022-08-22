package httplogger

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type key int

const (
	//LoggerCtxKey - This represents the logger key
	LoggerCtxKey key = 0
)

//SetLoggerLevelFromConfig - This API sets the log levelContext
func SetLoggerLevelFromConfig(logLevel string) log.Level {
	if logLevel != "" {
		level, err := log.ParseLevel(logLevel)
		if err != nil {
			println("Error while setting log level: " + logLevel + ". Root cause: " + err.Error())
		} else {
			return level
		}
	}
	return log.InfoLevel
}

func GetLogger(ctx context.Context) *log.Entry {
	raw := ctx.Value(LoggerCtxKey)
	if raw == nil {
		raw = log.WithFields(log.Fields{})
	}
	return raw.(*log.Entry)
}
