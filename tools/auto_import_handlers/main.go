package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	src := flag.String("src", "./internal/handlers", "Путь к папке c хэндлерами")
	dst := flag.String("dst", "./cmd/site/import_handlers.go", "Куда добавить импорт хэндлеров")
	pkg := flag.String("pkg", "main", "Имя package в dst файле")
	module := flag.String("module", "personal-site", "Go модуль")
	flag.Parse()

	handlerPackages, err := findHandlers(*src, *module)
	if err != nil {
		log.Fatalf("Ошибка при обходе папки %s: %v", *src, err)
	}

	err = generateFile(*dst, *pkg, handlerPackages)
	if err != nil {
		log.Fatalf("Не удалось сгенерировать файл %s: %v", *dst, err)
	}

	fmt.Printf("Сгенерирован файл: %s\n", *dst)
}

// findHandlers Проходит по папке src, ищет папки с .go и формирует список для импорта
func findHandlers(srcDir, moduleName string) ([]string, error) {
	var packagesName []string

	absSrc, err := filepath.Abs(srcDir)
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(absSrc, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() || path == absSrc {
			return nil
		}

		goFiles, _ := filepath.Glob(filepath.Join(path, "*.go"))
		if len(goFiles) == 0 {
			return nil
		}

		relPath, err := filepath.Rel(absSrc, path)
		if err != nil {
			return err
		}

		relSrcDir := strings.TrimPrefix(srcDir, "./")

		importPath := filepath.ToSlash(filepath.Join(moduleName, relSrcDir, relPath))

		packagesName = append(packagesName, importPath)

		return nil
	})

	if err != nil {
		return nil, err
	}
	return packagesName, nil
}

// generateFile Непосредственно записывает в dst файл необходимые импорты
func generateFile(dst, pkg string, importPaths []string) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("package %s\n\n", pkg))
	sb.WriteString("import (\n")
	for _, ip := range importPaths {
		sb.WriteString(fmt.Sprintf("\t_ \"%s\"\n", ip))
	}
	sb.WriteString(")\n")

	return ioutil.WriteFile(dst, []byte(sb.String()), 0644)
}
