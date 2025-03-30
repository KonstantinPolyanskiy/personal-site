package main

import "personal-site/internal/logging"

//go:generate go run ../../tools/auto_import_handlers/main.go -src=../../internal/handlers -dst=./import_handlers.go -pkg=main -module=personal-site
func main() {
	loggerRegistry := logging.NewRegistry(nil)

}
