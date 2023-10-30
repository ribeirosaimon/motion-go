package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func FindRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho atual:", err)
		return "", fmt.Errorf("raiz do projeto não encontrada")
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)

		if parentDir == currentDir {
			break
		}

		currentDir = parentDir
	}

	return "", fmt.Errorf("raiz do projeto não encontrada")
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if re.MatchString(email) {
		return true
	}
	return false
}
