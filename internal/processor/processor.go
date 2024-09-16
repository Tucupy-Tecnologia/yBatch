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
	"github.com/nfnt/resize"
)

func ProcessImages(inputPath, outputDir string, quality float32, lossless bool, losslessCompression int, width, height int) error {
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
			if err := processImage(file, outputDir, quality, lossless, losslessCompression, width, height); err != nil {
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

func processImage(inputPath string, outputDir string, quality float32, lossless bool, losslessCompression int, width, height int) error {
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		return fmt.Errorf("error getting file info: %v", err)
	}

	fileName := fileInfo.Name()
	if fileInfo.IsDir() || strings.HasPrefix(fileName, ".") {
		return nil // Skip directories and dotfiles
	}

	img, err := loadImage(inputPath)
	if err != nil {
		return err
	}

	if width > 0 && height > 0 {
		img = resizeImage(img, width, height)
	}

	outputPath := filepath.Join(outputDir, strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))+".webp")

	return saveAsWebP(img, outputPath, quality, lossless, losslessCompression)
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("error decoding image %s: %v", path, err)
	}

	return img, nil
}

func resizeImage(img image.Image, width, height int) image.Image {
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}

func saveAsWebP(img image.Image, outputPath string, quality float32, lossless bool, losslessCompression int) error {
	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer output.Close()

	var options *encoder.Options
	if lossless {
		options, err = encoder.NewLosslessEncoderOptions(encoder.PresetDefault, losslessCompression)
	} else {
		options, err = encoder.NewLossyEncoderOptions(encoder.PresetDefault, quality)
	}
	if err != nil {
		return fmt.Errorf("error creating encoder options: %v", err)
	}

	if err := webp.Encode(output, img, options); err != nil {
		return fmt.Errorf("error encoding to WebP: %v", err)
	}

	return nil
}
