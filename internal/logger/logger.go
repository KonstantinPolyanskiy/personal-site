package logger

import "go.uber.org/zap"

// CommonSettings Настройки, которые есть у каждого логгера модуля
type CommonSettings struct {
	Enabled              bool
	WriteToMainLogFile   bool
	WriteToModuleLogFile bool
	ExternalUrl          string
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
