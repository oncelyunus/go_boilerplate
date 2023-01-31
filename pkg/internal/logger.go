package internal

import (
	"fmt"
	"os"

	"github.com/oncelyunus/go_boilerplate/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type AppLogger struct {
	config *config.Config
	Logger *zap.SugaredLogger
}

func NewLogger(cfg *config.Config) *AppLogger {
	return &AppLogger{
		config: cfg,
	}
}

func (app *AppLogger) Init() *zap.SugaredLogger {
	var filename = fmt.Sprintf("%s.log", app.config.Server.Application)
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	/*logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile) */
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    app.config.Logger.MaxSize, // megabytes
		MaxBackups: app.config.Logger.MaxBackups,
		MaxAge:     app.config.Logger.MaxAge, // days
		Compress:   app.config.Logger.Compress,
	})

	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger.Sugar()
}
