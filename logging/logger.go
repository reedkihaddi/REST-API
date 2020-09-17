package logging

import (
	"log"
	"go.uber.org/zap"
)

// Log is used to access log within other pacakages.
var Log *zap.SugaredLogger

//InitLogger to initialize the logger.
//TODO create a log file and add log rolling and other logging related stuff.
func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Cannot start logging.")
	}
	Log = logger.Sugar()
}
