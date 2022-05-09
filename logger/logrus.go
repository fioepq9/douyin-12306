package logger

import (
	"douyin-12306/config"
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"os"
)

func newLogrusLogger() *logrusLogger {
	l := logrus.New()

	// set log-level
	switch config.C.Log.Level {
	case "panic":
		l.SetLevel(logrus.PanicLevel)
	case "fatal":
		l.SetLevel(logrus.FatalLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	case "warn":
		l.SetLevel(logrus.WarnLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "trace":
		l.SetLevel(logrus.TraceLevel)
	default:
		panic(fmt.Errorf(
			"the log level must be (panic fatal error warn info debug trace), but is %s",
			config.C.Log.Level))
	}

	// set output
	if config.C.Log.Out != "stdout" {
		output, err := os.OpenFile(config.C.Log.Out, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic("log init error " + err.Error())
		}
		l.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006/01/02-15:04:05"})
		l.SetOutput(output)
	} else {
		// use go-colorable
		l.SetFormatter(&logrus.TextFormatter{ForceColors: true, TimestampFormat: "2006/01/02-15:04:05"})
		l.SetOutput(colorable.NewColorableStdout())
	}

	return &logrusLogger{logger: l}
}

type logrusLogger struct {
	logger *logrus.Logger
}

func (l *logrusLogger) Panic(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Panic(msg)
}

func (l *logrusLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(msg)
}

func (l *logrusLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(msg)
}

func (l *logrusLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(msg)
}

func (l *logrusLogger) Info(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Info(msg)
}

func (l *logrusLogger) Debug(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Debug(msg)
}

func (l *logrusLogger) Trace(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Trace(msg)
}
