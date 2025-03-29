package logging

import (
	"encoding/json"
	"go.uber.org/zap/zapcore"
	"personal-site/pkg/common_data"
)

func (cls *CommonLoggerSettings) Settings() {}

// CommonLoggerSettings Настройки, которые есть у каждого логгера модуля
type CommonLoggerSettings struct {
	Enabled              bool
	EnabledLevels        []string
	WriteToMainLogFile   bool
	WriteToModuleLogFile bool
	WriteToConsole       bool
	ExternalUrl          string

	Identity common_data.IdentityInformation
	Creation common_data.CreationInformation
}

// ToJson Сериализует структуру в Json
func (cls *CommonLoggerSettings) ToJson() ([]byte, error) {
	data, err := json.Marshal(cls)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cls *CommonLoggerSettings) Name() string {
	return *cls.Identity.Name
}

func (cls *CommonLoggerSettings) IsLevelEnabled(level zapcore.Level) bool {
	for _, lvl := range cls.EnabledLevels {
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
