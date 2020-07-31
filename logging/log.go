package logging

import "go.uber.org/zap/zapcore"

// Interface logging interface
type Interface interface {
	SetLevel(zapcore.Level)

	Debug(args ...interface{})
	Debugf(fmt string, args ...interface{})
	Info(args ...interface{})
	Infof(fmt string, args ...interface{})
	Warn(args ...interface{})
	Warnf(fmt string, args ...interface{})
	Error(args ...interface{})
	Errorf(fmt string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

// Utility functions
func SetLevel(level zapcore.Level)              { defaultLogger.SetLevel(level) }
func Info(args ...interface{})                  { defaultLogger.Info(args...) }
func Infof(format string, args ...interface{})  { defaultLogger.Infof(format, args...) }
func Debug(args ...interface{})                 { defaultLogger.Debug(args...) }
func Debugf(format string, args ...interface{}) { defaultLogger.Debugf(format, args...) }
func Warn(args ...interface{})                  { defaultLogger.Warn(args...) }
func Warnf(format string, args ...interface{})  { defaultLogger.Warnf(format, args...) }
func Error(args ...interface{})                 { defaultLogger.Error(args...) }
func Errorf(format string, args ...interface{}) { defaultLogger.Errorf(format, args...) }
func Fatal(args ...interface{})                 { defaultLogger.Fatal(args...) }
func Fatalf(format string, args ...interface{}) { defaultLogger.Fatalf(format, args...) }
