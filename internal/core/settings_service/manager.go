package settings_service

import (
	"errors"
	"go.uber.org/zap"
	"personal-site/internal/logging"
	"reflect"
)

// SimpleSettingStorageItem Описывает, может ли структура быть сохранена в БД как Json
type SimpleSettingStorageItem interface {
	ToJson() ([]byte, error)
	Name() string
}

// Settings Описывает более общую сущность "Настройка"
type Settings interface {
	Settings()
}

// DaoRegistry DAO для работы с простыми настройками
type DaoRegistry interface {
	SimpleSettingsDao() SimpleSettingsDao
}

// SimpleSettingsDao для работы с простыми настройками
type SimpleSettingsDao interface {
	Create(name string, jsonData []byte) error
}

type SettingsService struct {
	simpleSettingsDao SimpleSettingsDao
	logger            *logging.ModuleLogger
}

func (sm *SettingsService) Name() string {
	return "SettingsService"
}

func (sm *SettingsService) FirstInitialization(settings Settings) (Settings, error) {
	if simpleSettings, ok := settings.(SimpleSettingStorageItem); ok {
		json, err := simpleSettings.ToJson()
		if err != nil {
			sm.logger.Error("Failed to serialize simple settings", zap.Error(err))
			return nil, err
		}

		err = sm.simpleSettingsDao.Create(simpleSettings.Name(), json)
		if err != nil {
			return nil, err
		}

		sm.logger.Info("Created simple settings", zap.String("simpleSettings", reflect.TypeOf(simpleSettings).Name()))
		return settings, nil
	} else {
		return nil, errors.New("not implemented other settings")
	}
}

func New(lr *logging.Registry, dr DaoRegistry) *SettingsService {
	return &SettingsService{
		logger: lr.LoggerFor("SettingsManagerLogger"),
		//simpleSettingsDao: dr.SimpleSettingsDao(),
	}
}
