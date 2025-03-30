package service_locator

import (
	"go.uber.org/zap"
	"personal-site/internal/logging"
	"sync"
)

type ServiceLocator struct {
	mu sync.RWMutex

	services map[string]interface{}
	logger   *logging.ModuleLogger
}

func New(lr *logging.Registry) *ServiceLocator {
	return &ServiceLocator{
		services: make(map[string]interface{}),
		logger:   lr.LoggerFor("ServiceLocatorLogger"),
	}
}

func (sl *ServiceLocator) Register(name string, service interface{}) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	sl.services[name] = service
	sl.logger.Info("Service registered", zap.String("service", name))
}

func (sl *ServiceLocator) Get(name string) (interface{}, bool) {
	sl.mu.RLock()
	defer sl.mu.RUnlock()

	service, exist := sl.services[name]
	if !exist {
		sl.logger.Warn("Service not found", zap.String("requested service", name))
	}

	return service, exist
}
