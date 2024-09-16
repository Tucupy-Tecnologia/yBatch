package processor

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func ParseAsWebp(inputPath, outputDir string, quality float32, lossless bool, losslessCompression int) error {
	if outputDir == "" {
		inputAbsPath, err := filepath.Abs(inputPath)
		if err != nil {
			return fmt.Errorf("error getting absolute path: %v", err)
		}
		parentDir := filepath.Dir(inputAbsPath)
		outputDir = filepath.Join(parentDir, "yBatch-output-"+filepath.Base(inputPath))
	}

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory: %v", err)
	}

	files, err := filepath.Glob(filepath.Join(inputPath, "*"))
	if err != nil {
		return fmt.Errorf("error getting input files: %v", err)
	}

	var wg sync.WaitGroup
	errors := make(chan error, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if err := convertToWebP(file, outputDir, quality, lossless, losslessCompression); err != nil {
				errors <- fmt.Errorf("error processing %s: %v", file, err)
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		fmt.Println(err)
	}

	return nil
}

func convertToWebP(inputPath string, outputDir string, quality float32, lossless bool, losslessCompression int) error {
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		return fmt.Errorf("error getting file info: %v", err)
	}
	fileName := fileInfo.Name()
	if fileInfo.IsDir() || strings.HasPrefix(fileName, ".") {
		return nil // Skip directories and dotfiles
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding image %s: %v", inputPath, err)
	}

	// Remove the old extension and add .webp
	baseFileName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	outputPath := filepath.Join(outputDir, baseFileName+".webp")

	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer output.Close()

	var options *encoder.Options

	if lossless {
		options, err = encoder.NewLosslessEncoderOptions(encoder.PresetDefault, losslessCompression)
		if err != nil {
			return fmt.Errorf("error creating lossless encoder options: %v", err)
		}
	} else {
		options, err = encoder.NewLossyEncoderOptions(encoder.PresetDefault, quality)
		if err != nil {
			return fmt.Errorf("error creating lossy encoder options: %v", err)
		}
	}

	if err := webp.Encode(output, img, options); err != nil {
		return fmt.Errorf("error encoding to WebP: %v", err)
	}

	return nil
}
