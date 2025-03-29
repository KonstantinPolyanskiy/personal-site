package logging

import "go.uber.org/zap"

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
