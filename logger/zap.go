package logger

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewZapLogger(out, levelStr string) *zapLogger {
	encoder := newEncoder(out)
	writeSyncer, err := newWriteSyncer(out)
	if err != nil {
		panic(errors.Wrap(err, "zap logger init"))
	}
	level, err := newLogLevel(levelStr)
	if err != nil {
		panic(errors.Wrap(err, "zap logger init"))
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	return &zapLogger{
		logger: zap.New(
			core,
			zap.AddStacktrace(zap.PanicLevel),
			zap.AddCaller(),
			zap.AddCallerSkip(1)),
	}
}

func newEncoder(out string) (encoder zapcore.Encoder) {
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	if out == "stdout" {
		encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConf)
	} else {
		encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder = zapcore.NewJSONEncoder(encoderConf)
	}
	return encoder
}

func newWriteSyncer(out string) (writeSyncer zapcore.WriteSyncer, err error) {
	if out == "stdout" {
		writeSyncer = zapcore.AddSync(os.Stdout)
	} else {
		output, err := os.OpenFile(out, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		writeSyncer = zapcore.AddSync(output)
	}
	return writeSyncer, nil
}

func newLogLevel(levelStr string) (level zapcore.Level, err error) {
	switch levelStr {
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "error":
		level = zapcore.ErrorLevel
	case "warn":
		level = zapcore.WarnLevel
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	case "trace":
		level = zapcore.DebugLevel
	default:
		return 0, fmt.Errorf("the log level must be (panic fatal error warn info debug trace), but is %s", levelStr)
	}
	return level, nil
}

type zapLogger struct {
	logger *zap.Logger
}

func (z *zapLogger) Sync() error {
	return z.logger.Sugar().Sync()
}

func (z *zapLogger) Panic(args ...interface{}) {
	z.logger.Sugar().Panic(args...)
}

func (z *zapLogger) Panicf(format string, args ...interface{}) {
	z.logger.Sugar().Panicf(format, args...)
}

func (z *zapLogger) Panicw(msg string, fields map[string]interface{}) {
	keyAndValues := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyAndValues = append(keyAndValues, key)
		keyAndValues = append(keyAndValues, value)
	}
	z.logger.Sugar().Panicw(msg, keyAndValues...)
}

func (z *zapLogger) Fatal(args ...interface{}) {
	z.logger.Sugar().Fatal(args...)
}

func (z *zapLogger) Fatalf(format string, args ...interface{}) {
	z.logger.Sugar().Fatalf(format, args...)
}

func (z *zapLogger) Fatalw(msg string, fields map[string]interface{}) {
	keyAndValues := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyAndValues = append(keyAndValues, key)
		keyAndValues = append(keyAndValues, value)
	}
	z.logger.Sugar().Fatalw(msg, keyAndValues...)
}

func (z *zapLogger) Error(args ...interface{}) {
	z.logger.Sugar().Error(args...)
}

func (z *zapLogger) Errorf(format string, args ...interface{}) {
	z.logger.Sugar().Errorf(format, args...)
}

func (z *zapLogger) Errorw(msg string, fields map[string]interface{}) {
	keyAndValues := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyAndValues = append(keyAndValues, key)
		keyAndValues = append(keyAndValues, value)
	}
	z.logger.Sugar().Errorw(msg, keyAndValues...)
}

func (z *zapLogger) Warn(args ...interface{}) {
	z.logger.Sugar().Warn(args...)
}

func (z *zapLogger) Warnf(format string, args ...interface{}) {
	z.logger.Sugar().Warnf(format, args...)
}

func (z *zapLogger) Warnw(msg string, fields map[string]interface{}) {
	keyAndValues := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyAndValues = append(keyAndValues, key)
		keyAndValues = append(keyAndValues, value)
	}
	z.logger.Sugar().Warnw(msg, keyAndValues...)
}

func (z *zapLogger) Info(args ...interface{}) {
	z.logger.Sugar().Info(args...)
}

func (z *zapLogger) Infof(format string, args ...interface{}) {
	z.logger.Sugar().Infof(format, args...)
}

func (z *zapLogger) Infow(msg string, fields map[string]interface{}) {
	keyAndValues := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyAndValues = append(keyAndValues, key)
		keyAndValues = append(keyAndValues, value)
	}
	z.logger.Sugar().Infow(msg, keyAndValues...)
}

func (z *zapLogger) Debug(args ...interface{}) {
	z.logger.Sugar().Debug(args...)
}

func (z *zapLogger) Debugf(format string, args ...interface{}) {
	z.logger.Sugar().Debugf(format, args...)
}

func (z *zapLogger) Debugw(msg string, fields map[string]interface{}) {
	keyAndValues := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyAndValues = append(keyAndValues, key)
		keyAndValues = append(keyAndValues, value)
	}
	z.logger.Sugar().Debugw(msg, keyAndValues...)
}

func (z *zapLogger) Trace(args ...interface{}) {
	z.logger.Sugar().Debug(args...)
}

func (z *zapLogger) Tracef(format string, args ...interface{}) {
	z.logger.Sugar().Debugf(format, args...)
}

func (z *zapLogger) Tracew(msg string, fields map[string]interface{}) {
	keyAndValues := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyAndValues = append(keyAndValues, key)
		keyAndValues = append(keyAndValues, value)
	}
	z.logger.Sugar().Debugw(msg, keyAndValues...)
}
