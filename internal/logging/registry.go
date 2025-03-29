package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Registry struct {
	mainLogger *zap.Logger

	// Ключ - название модуля
	settings map[string]CommonLoggerSettings
	loggers  map[string]*zap.Logger

	mu sync.RWMutex
}

func NewRegistry(mainLogger *zap.Logger) *Registry {
	return &Registry{
		mainLogger: mainLogger,
		settings:   make(map[string]CommonLoggerSettings),
		loggers:    make(map[string]*zap.Logger),
	}
}

func (lr *Registry) LoggerFor(moduleLoggerName string) *ModuleLogger {
	return &ModuleLogger{
		registry:   lr,
		moduleName: moduleLoggerName,
	}
}

func (lr *Registry) log(moduleLoggerName, customPath string, level zapcore.Level, msg string, fields ...zap.Field) {
	lr.mu.RLock()
	settings, exists := lr.settings[moduleLoggerName]
	lr.mu.RUnlock()

	if !exists || !settings.Enabled || !settings.IsLevelEnabled(level) {
		return
	}

	lr.mu.RLock()
	logger, exists := lr.loggers[moduleLoggerName]
	lr.mu.RUnlock()

	if !exists {
		path := fmt.Sprintf("logs/%s", strings.TrimSuffix(moduleLoggerName, "Logger"))
		if customPath != "" {
			path = customPath
		}
		logger = lr.createModuleLogger(moduleLoggerName, path, settings)
	}

	logger.Check(level, msg).Write(fields...)
}

func (lr *Registry) createModuleLogger(moduleLoggerName, path string, settings CommonLoggerSettings) *zap.Logger {
	var cores []zapcore.Core

	if settings.WriteToMainLogFile {
		cores = append(cores, lr.mainLogger.Core())
	}

	if settings.WriteToModuleLogFile {
		logPath := filepath.Join(path, fmt.Sprintf("%s.log", strings.ToLower(strings.TrimSuffix(moduleLoggerName, "Logger"))))
		if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
			lr.mainLogger.Error("failed to create log directory", zap.String("path", logPath), zap.Error(err))
		}

		ws := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    100,
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   true,
		})

		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), ws, zap.InfoLevel)
		cores = append(cores, core)
	}

	var logger *zap.Logger
	if len(cores) == 0 {
		logger = zap.NewNop()
	} else {
		logger = zap.New(zapcore.NewTee(cores...)).Named(moduleLoggerName)
	}

	lr.mu.Lock()
	lr.loggers[moduleLoggerName] = logger
	lr.mu.Unlock()

	lr.mainLogger.Info("Initialized new module logging", zap.String("module", moduleLoggerName), zap.String("path", path))

	return logger
}

func (lr *Registry) UpdateConfig(moduleLoggerName string, newSettings CommonLoggerSettings) {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	lr.settings[moduleLoggerName] = newSettings
	delete(lr.loggers, moduleLoggerName) // чтобы перестроился на следующем вызове

	lr.mainLogger.Info("Logger settings updated",
		zap.String("module", moduleLoggerName),
		zap.Any("new settings", newSettings))
}
