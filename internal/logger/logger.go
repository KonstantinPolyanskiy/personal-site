package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CommonSettings Настройки, которые есть у каждого логгера модуля
type CommonSettings struct {
	Enabled              bool
	EnabledLevels        []string
	WriteToMainLogFile   bool
	WriteToModuleLogFile bool
	WriteToConsole       bool
	ExternalUrl          string
}

func (s *CommonSettings) IsLevelEnabled(level zapcore.Level) bool {
	for _, lvl := range s.EnabledLevels {
		if lvl == level.String() {
			return true
		}
	}

	return false
}

type ModuleLogger struct {
	registry *Registry

	moduleName string
	customPath string // Для нетипичного пути к файлу логов модуля
}

func (ml *ModuleLogger) UseAtypicalPath(path string) *ModuleLogger {
	ml.customPath = path

	return ml
}

func (ml *ModuleLogger) Info(msg string, fields ...zap.Field) {
	ml.registry.log(ml.moduleName, ml.customPath, zap.InfoLevel, msg, fields...)
	ml.customPath = ""
}

func (ml *ModuleLogger) Debug(msg string, fields ...zap.Field) {
	ml.registry.log(ml.moduleName, ml.customPath, zap.DebugLevel, msg, fields...)
	ml.customPath = ""
}

func (ml *ModuleLogger) Error(msg string, fields ...zap.Field) {
	ml.registry.log(ml.moduleName, ml.customPath, zap.ErrorLevel, msg, fields...)
	ml.customPath = ""
}

func (ml *ModuleLogger) Fatal(msg string, fields ...zap.Field) {
	ml.registry.log(ml.moduleName, ml.customPath, zap.FatalLevel, msg, fields...)
	ml.customPath = ""
}

func (ml *ModuleLogger) Warn(msg string, fields ...zap.Field) {
	ml.registry.log(ml.moduleName, ml.customPath, zap.WarnLevel, msg, fields...)
	ml.customPath = ""
}
