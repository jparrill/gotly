package config

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var MainLogger *zap.Logger

// InitLogger function initializes the Logger engine using Zap as a base
func InitLogger() {
	_, err := os.Stat(CFG.TMPFolder)
	if err != nil {
		fmt.Println("BasePath for Log entries does not exists, creating...")
		os.Mkdir(CFG.TMPFolder, 0754)
	}

	var defaultLogLevel zapcore.LevelEnabler

	loggerConfig := zap.NewProductionEncoderConfig()
	loggerConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(loggerConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(loggerConfig)

	var logFile *os.File

	if CFG.Log.LogTruncate {
		logFile, err = os.OpenFile(CFG.Log.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Panicf("Error openning the %s logfile to store logs\n", CFG.Log.LogPath)
		}
	} else {
		logFile, err = os.OpenFile(CFG.Log.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Panicf("Error openning the %s logfile to store logs\n", CFG.Log.LogPath)
		}
	}
	writer := zapcore.AddSync(logFile)

	switch CFG.Log.LogLevel {
	case "debug":
		defaultLogLevel = zapcore.DebugLevel
	case "info":
		defaultLogLevel = zapcore.InfoLevel
	case "warn":
		defaultLogLevel = zapcore.WarnLevel
	case "error":
		defaultLogLevel = zapcore.ErrorLevel
	case "panic":
		defaultLogLevel = zapcore.PanicLevel
	case "fatal":
		defaultLogLevel = zapcore.FatalLevel
	default:
		log.Panicln(`LogLevel Not implemented, please check "go doc zapcore.DebugLevel"`)
	}
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	MainLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	defer MainLogger.Sync()

	MainLogger.Info("Logger construction succeeded")
}
