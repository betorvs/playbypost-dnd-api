package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"go.uber.org/zap/zaptest/observer"
)

var (
	// LoggerMasterCalls int
	LoggerMasterCalls int
	// LoggerErrorMasterCalls int
	LoggerErrorMasterCalls int
	// LoggerPlayerCalls int
	LoggerPlayerCalls int
	// LoggerErrorPlayerCalls int
	LoggerErrorPlayerCalls int
)

//MockLogger wrapper with two Logger methods
type MockLogger struct {
	//Logger only supports structured logging (less features more performance)
	Logger *zap.Logger
	//Sugar supports structured and printf-style APIs (less performance more features)
	Sugar *zap.SugaredLogger
	// Observer to access logs in tests
	Observer *observer.ObservedLogs
}

// InitMockLogger func creates a mock logger interface
func InitMockLogger() appcontext.Component {
	t := new(testing.T)
	loggerMock := zaptest.NewLogger(t)
	// loggerMock := zaptest.NewLogger(t, zaptest.WrapOptions(zap.Hooks(func(e zapcore.Entry) error {
	// 	// if e.Level == zap.ErrorLevel {
	// 	// 	t.Fatal("Error should never happen!")
	// 	// }
	// 	return nil
	// })))
	sugar := loggerMock.Sugar()
	_, observer := observer.New(zap.InfoLevel)
	return MockLogger{Logger: loggerMock, Sugar: sugar, Observer: observer}
	// return MockLogger{}
}

// GetLogs func mock
func (logger MockLogger) GetLogs() []observer.LoggedEntry {
	return logger.Observer.All()
}

// Debug uses fmt.Sprint to construct and log a message.
func (logger MockLogger) Debug(args ...interface{}) {
	// logger.Sugar.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func (logger MockLogger) Info(args ...interface{}) {
	// logger.Sugar.Info(args)
	text := fmt.Sprintf("%v", args...)
	// logger.Logger.Info(text)
	// fmt.Println(text)
	// logs := logger.GetLogs()
	// if len(logs) != 0 {
	if strings.Contains(text, "Master message processed") {
		LoggerMasterCalls++
	}
	if strings.Contains(text, "Master request not processed") {
		LoggerErrorMasterCalls++
	}
	if strings.Contains(text, "Player message processed") {
		LoggerPlayerCalls++
	}
	if strings.Contains(text, "Player request not processed") {
		LoggerErrorPlayerCalls++
	}
	// }

	// logger.Sugar.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func (logger MockLogger) Warn(args ...interface{}) {
	// logger.Sugar.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func (logger MockLogger) Error(args ...interface{}) {
	// logger.Sugar.Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (logger MockLogger) DPanic(args ...interface{}) {
	// logger.Sugar.DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (logger MockLogger) Panic(args ...interface{}) {
	// logger.Sugar.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (logger MockLogger) Fatal(args ...interface{}) {
	// logger.Sugar.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (logger MockLogger) Debugf(template string, args ...interface{}) {
	// logger.Sugar.Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func (logger MockLogger) Infof(template string, args ...interface{}) {
	// logger.Sugar.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (logger MockLogger) Warnf(template string, args ...interface{}) {
	// logger.Sugar.Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (logger MockLogger) Errorf(template string, args ...interface{}) {
	// logger.Sugar.Errorf(template, args)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (logger MockLogger) DPanicf(template string, args ...interface{}) {
	// logger.Sugar.DPanicf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (logger MockLogger) Panicf(template string, args ...interface{}) {
	// logger.Sugar.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (logger MockLogger) Fatalf(template string, args ...interface{}) {
	// logger.Sugar.Fatalf(template, args)
}

//Sync flushes the log if needed
func (logger MockLogger) Sync() {
	// _ = logger.Sugar.Sync()
}
