package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// init setup logger
func Setup() {
	// Define log level
	level := zap.NewAtomicLevel()
	level.SetLevel(zap.DebugLevel)

	// Create new zap logger
	logger = zap.New(zapcore.NewCore(
		getEncoder(),
		getLogWriter(),
		level,
	))
	defer logger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderCfg)
	// return zapcore.NewJSONEncoder(encoderCfg)
}

func getLogWriter() zapcore.WriteSyncer {
	// file, _ := os.Create("./test.log")
	// return zapcore.AddSync(file)
	return zapcore.Lock(os.Stdout)
}

// Info log
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Debug log
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Error log
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Fatal log
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
