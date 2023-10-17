// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	ierr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/errors"
)

type Config struct {
	LogLevel       string
	ServiceName    string
	ServiceVersion string
}

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Panic(args ...any)
	Panicf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Sync() error
}

type logger struct {
	logger *zap.SugaredLogger
}

func Default(serviceName string) Logger {
	return NewLogger(&Config{
		LogLevel:    "debug",
		ServiceName: serviceName,
	})
}

func NewLogger(cfg *Config) Logger {
	return &logger{
		logger: initLogger(cfg),
	}
}

func initLogger(cfg *Config) *zap.SugaredLogger {
	zap.NewProduction()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	encoder := zapcore.NewJSONEncoder(encoderCfg)
	logWriter := zapcore.AddSync(os.Stderr)
	logLevel := getLoggerLevel(cfg.LogLevel)
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))

	fields := getFields(cfg)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Fields(fields...))

	return logger.Sugar()
}

func getLoggerLevel(level string) zapcore.Level {
	switch level {
	case "debug", "DEBUG":
		return zap.DebugLevel
	case "info", "INFO":
		return zap.InfoLevel
	case "warn", "WARN":
		return zap.WarnLevel
	case "error", "ERROR":
		return zap.ErrorLevel
	case "panic", "PANIC":
		return zap.PanicLevel
	case "fatal", "FATAL":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func getFields(cfg *Config) []zapcore.Field {
	return []zapcore.Field{
		{
			Key:    "svc",
			Type:   zapcore.StringType,
			String: cfg.ServiceName,
		},
		{
			Key:    "ver",
			Type:   zapcore.StringType,
			String: cfg.ServiceVersion,
		},
		{
			Key:    "type",
			Type:   zapcore.StringType,
			String: "log",
		},
	}
}

func (o logger) Debug(args ...any) {
	o.logger.Debug(args...)
}

func (o logger) Debugf(format string, args ...any) {
	o.logger.Debugf(format, args...)
}

func (o logger) Info(args ...any) {
	o.logger.Info(args...)
}

func (o logger) Infof(format string, args ...any) {
	o.logger.Infof(format, args...)
}

func (o logger) Warn(args ...any) {
	o.logger.Warn(args...)
}

func (o logger) Warnf(format string, args ...any) {
	o.logger.Warnf(format, args...)
}

func (o logger) Error(args ...any) {
	o.logger.Error(args...)
}

func (o logger) Errorf(format string, args ...any) {
	o.logger.Errorf(format, args...)
}

func (o logger) Panic(args ...any) {
	o.logger.Panic(args...)
}

func (o logger) Panicf(format string, args ...any) {
	o.logger.Panicf(format, args...)
}

func (o logger) Fatal(args ...any) {
	o.logger.Fatal(args...)
}

func (o logger) Fatalf(format string, args ...any) {
	o.logger.Fatalf(format, args...)
}

func (o logger) Sync() error {
	if err := o.logger.Sync(); err != nil {
		return ierr.WrapErrorf(err, ierr.Unknown, "logger.Sync")
	}

	return nil
}
