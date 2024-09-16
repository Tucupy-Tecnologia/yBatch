package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Tucupy-Tecnologia/yBatch/internal"
)

func ValidateDirectory(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error accessing the path: %v\n", err)
		os.Exit(1)
	}
	if !fileInfo.IsDir() {
		fmt.Println("The provided path is not a directory")
		os.Exit(1)
	}
}

func ValidateImageFiles(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	validFiles := 0
	for _, file := range files {
		// Skip dot files
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if file.IsDir() {
			fmt.Println("The directory contains subdirectories")
			os.Exit(1)
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if !internal.SupportedExtensions[ext] {
			fmt.Printf("The directory contains non-image file: %s\n", file.Name())
			os.Exit(1)
		}

		validFiles++
	}

	if validFiles == 0 {
		fmt.Println("The directory does not contain any valid image files")
		os.Exit(1)
	}
}
