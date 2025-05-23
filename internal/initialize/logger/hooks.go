package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Hooks 日志钩子
func Hooks() zap.Option {
	return zap.Hooks(func(entry zapcore.Entry) error {
		if entry.Level == zapcore.ErrorLevel {
			// todo
		}
		return nil
	})
}
