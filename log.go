// log.go
package log

import "github.com/astaxie/beego/logs"

var Log = logs.NewLogger(10000)

func init() {
	Log.SetLogger("console", "")
	Log.EnableFuncCallDepth(true)
	Log.SetLogFuncCallDepth(2)
	Log.SetLevel(logs.LevelDebug)

}

//	LevelAlert
//	LevelCritical
//	LevelError
//	LevelWarning
//	LevelNotice
//	LevelInformational
//	LevelDebug
func Debug(v ...interface{}) {
	Log.Debug("%v", v)
}

func Info(v ...interface{}) {
	Log.Info("%v", v)
}

func Notice(v ...interface{}) {
	Log.Notice("%v", v)
}

func Warning(v ...interface{}) {
	Log.Warning("%v", v)
}

func Error(v ...interface{}) {
	Log.Error("%v", v)
}

func Critical(v ...interface{}) {
	Log.Critical("%v", v)
}

func Alert(v ...interface{}) {
	Log.Alert("%v", v)
}
