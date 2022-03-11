package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Log.SetLevel(WARN)
	Log.SetPrefix("test prefix")
	Log.Info("info")
	Log.Debug("debug")
	Log.Error("error")
	Log.Warn("warn")
	//Log.Panic("panic")
	//Log.Fatal("fatal")
	Error("error")
	Debug("debug")
	Warn("warn")
	Fatal("fatal")
}
