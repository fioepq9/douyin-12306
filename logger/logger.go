package logger

import (
	"douyin-12306/config"
)

var L logger

type logger interface {
	Sync() error

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicw(msg string, fields map[string]interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalw(msg string, fields map[string]interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorw(msg string, fields map[string]interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnw(msg string, fields map[string]interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infow(msg string, fields map[string]interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugw(msg string, fields map[string]interface{})

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Tracew(msg string, fields map[string]interface{})
}

func init() {
	L = NewZapLogger(config.C.Log.Out, config.C.Log.Level)
	L.Infow("Init logger success", map[string]interface{}{
		"Logger config": config.C.Log,
	})
}
