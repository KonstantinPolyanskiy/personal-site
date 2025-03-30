package main

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"personal-site/internal/core/router"
	"personal-site/internal/core/service_locator"
	"personal-site/internal/core/settings_service"
	"personal-site/internal/logging"
	"syscall"
	"time"
)

//go:generate go run ../../tools/auto_import_handlers/main.go -src=../../internal/handlers -dst=./import_handlers.go -pkg=main -module=personal-site
func main() {
	mainCtx, cancel := context.WithCancel(context.Background())

	mainLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	err = mainLogger.Sync()
	if err != nil {
		panic(err)
	}

	loggerRegistry := logging.NewRegistry(mainLogger)
	serviceLocator := service_locator.New(loggerRegistry)
	settingsService := settings_service.New(loggerRegistry, nil)

	registryRequiredService(serviceLocator, settingsService)

	r := router.New(loggerRegistry)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r.Init(),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c

		shutdownCtx, _ := context.WithTimeout(mainCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()

			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				mainLogger.Warn("Graceful shutdown timed out, force exit")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		cancel()
	}()

	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	<-mainCtx.Done()
}

type Service interface {
	Name() string
}

func registryRequiredService(locator *service_locator.ServiceLocator, services ...Service) {
	for _, service := range services {
		locator.Register(service.Name(), service)
	}
}
