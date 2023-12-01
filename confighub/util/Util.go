package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FindRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("project root not found")
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

	return "", fmt.Errorf("project root not found")
}
func FindModuleDir(module string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("project root not found")
	}
	index := strings.Index(currentDir, "/motion-go/")
	if index == -1 {
		return "", fmt.Errorf("'/motion-go/' not found in the path")
	}
	result := currentDir[index+len("/motion-go/"):]

	replace := strings.Replace(currentDir, result, module, -1)
	validatePath(replace)
	return replace, nil
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if re.MatchString(email) {
		return true
	}
	return false
}

func validatePath(directoryPath string) {
	cleanedPath := filepath.Clean(directoryPath)

	_, err := os.Stat(cleanedPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("The dir %s not exist.\n", cleanedPath)
		}
	}
}
