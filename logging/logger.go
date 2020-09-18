package logging

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Log is used to access log within other pacakages.
var Log *zap.SugaredLogger

//Log rolling with lumberjack package
var lumlog = &lumberjack.Logger{
	Filename:   "./logging/server.log",
	MaxSize:    1, // megabytes
	MaxBackups: 5,   // number of log files
	MaxAge:     1,   // days
}

//InitLogger to initialize the logger.
func InitLogger() {
	logger, err := zap.NewDevelopment(zap.Hooks(lumberjackZapHook))
	if err != nil {
		log.Fatal("Cannot start logging.")
	}
	Log = logger.Sugar()
}

func lumberjackZapHook(e zapcore.Entry) error {
	lumlog.Write([]byte(fmt.Sprintf("%+v\n", e)))
	return nil
}
