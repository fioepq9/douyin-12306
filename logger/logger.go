package logger

var L logger

type logger interface {
	Panic(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Debug(msg string, fields map[string]interface{})
	Trace(msg string, fields map[string]interface{})
}

func init() {
	L = newLogrusLogger()
	L.Info("logger init success", map[string]interface{}{
		"package":  "logger",
		"function": "init",
	})
}
