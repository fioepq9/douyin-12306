package repository

import (
	"context"
	"douyin-12306/config"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

func convertLogLevel(levelStr string) (level logger.LogLevel, err error) {
	switch levelStr {
	case "error":
		level = logger.Error
	case "warn":
		level = logger.Warn
	case "info":
		level = logger.Info
	case "silent":
		level = logger.Silent
	default:
		return 0, fmt.Errorf("undefined log level")
	}
	return level, nil
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
		output, err := os.OpenFile(config.C.Log.Out, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		writeSyncer = zapcore.AddSync(output)
	}
	return writeSyncer, nil
}

func newLogLevel(logLevel logger.LogLevel) (level zapcore.Level, err error) {
	switch logLevel {
	case logger.Error:
		level = zapcore.ErrorLevel
	case logger.Warn:
		level = zapcore.WarnLevel
	case logger.Info:
		level = zapcore.InfoLevel
	case logger.Silent:
		level = zapcore.DebugLevel
	default:
		return 0, fmt.Errorf("undefined log level")
	}
	return level, nil
}

func NewMySQLLogger(out, levelStr string, conf logger.Config) *mysqlLogger {
	encoder := newEncoder(out)
	writeSyncer, err := newWriteSyncer(out)
	if err != nil {
		panic(errors.Wrap(err, "MySQL logger init"))
	}
	logLevel, err := convertLogLevel(levelStr)
	if err != nil {
		panic(errors.Wrap(err, "MySQL logger init"))
	}
	level, err := newLogLevel(logLevel)
	if err != nil {
		panic(errors.Wrap(err, "MySQL logger init"))
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	return &mysqlLogger{
		encoder:     encoder,
		writeSyncer: writeSyncer,
		logLevel:    logLevel,
		conf:        conf,
		logger: zap.New(
			core,
			zap.AddStacktrace(zap.PanicLevel),
			zap.AddCaller(),
			zap.AddCallerSkip(3)),
	}
}

type mysqlLogger struct {
	encoder     zapcore.Encoder
	writeSyncer zapcore.WriteSyncer
	logLevel    logger.LogLevel
	conf        logger.Config
	logger      *zap.Logger
}

func (l *mysqlLogger) LogMode(logLevel logger.LogLevel) logger.Interface {
	level, err := newLogLevel(logLevel)
	if err != nil {
		l.Error(context.TODO(), err.Error())
		return nil
	}
	core := zapcore.NewCore(l.encoder, l.writeSyncer, level)
	return &mysqlLogger{
		encoder:     l.encoder,
		writeSyncer: l.writeSyncer,
		logLevel:    logLevel,
		conf:        l.conf,
		logger: zap.New(
			core,
			zap.AddStacktrace(zap.PanicLevel),
			zap.AddCaller(),
			zap.AddCallerSkip(3)),
	}
}

func (l *mysqlLogger) Info(ctx context.Context, s string, i ...interface{}) {
	l.logger.Sugar().Info(s, i)
}

func (l *mysqlLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.logger.Sugar().Warn(s, i)
}

func (l *mysqlLogger) Error(ctx context.Context, s string, i ...interface{}) {
	l.logger.Sugar().Error(s, i)
}

func (l *mysqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.logLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.logLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.conf.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.logger.Sugar().Errorw(err.Error(),
				"SpendTime[ms]", float64(elapsed.Nanoseconds())/1e6,
				"rows", "-",
				"sql", sql,
			)
		} else {
			l.logger.Sugar().Errorw(err.Error(),
				"SpendTime[ms]", float64(elapsed.Nanoseconds())/1e6,
				"rows", rows,
				"sql", sql,
			)
		}
	case elapsed > l.conf.SlowThreshold && l.conf.SlowThreshold != 0 && l.logLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.conf.SlowThreshold)
		if rows == -1 {
			l.logger.Sugar().Warnw(slowLog,
				"SpendTime[ms]", float64(elapsed.Nanoseconds())/1e6,
				"rows", "-",
				"sql", sql,
			)
		} else {
			l.logger.Sugar().Warnw(slowLog,
				"SpendTime[ms]", float64(elapsed.Nanoseconds())/1e6,
				"rows", rows,
				"sql", sql,
			)
		}
	case l.logLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.logger.Sugar().Infow("",
				"SpendTime[ms]", float64(elapsed.Nanoseconds())/1e6,
				"rows", "-",
				"sql", sql,
			)
			l.logger.Sugar().Info()
		} else {
			l.logger.Sugar().Infow("",
				"SpendTime[ms]", float64(elapsed.Nanoseconds())/1e6,
				"rows", rows,
				"sql", sql,
			)
		}
	}
}
