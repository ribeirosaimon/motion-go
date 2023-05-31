package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho atual:", err)
		return "", fmt.Errorf("raiz do projeto não encontrada")
	}

	for {
		// Verifica se existe um arquivo go.mod no diretório atual
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return currentDir, nil
		}

		// Vai para o diretório pai
		parentDir := filepath.Dir(currentDir)

		// Verifica se chegamos ao diretório raiz
		if parentDir == currentDir {
			break
		}

		// Atualiza o diretório atual para o diretório pai
		currentDir = parentDir
	}

	return "", fmt.Errorf("raiz do projeto não encontrada")
}
