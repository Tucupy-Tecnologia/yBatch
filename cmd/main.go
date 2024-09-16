package main

import (
	"fmt"
	"os"

	"github.com/Tucupy-Tecnologia/yBatch/internal/config"
	"github.com/Tucupy-Tecnologia/yBatch/internal/processor"
	"github.com/Tucupy-Tecnologia/yBatch/internal/utils"
)

func main() {
	config := config.ParseFlags()
	utils.ValidateDirectory(config.Path)
	utils.ValidateImageFiles(config.Path)
	fmt.Printf("The directory '%s' contains only image files. Ready to process.\n", config.Path)
	fmt.Printf("Processing with width: %d, height: %d, aspect ratio: %s, and output format: %s\n",
		config.Width, config.Height, config.AspectRatio, config.Format)

	//TODO: Make these configurable
	quality := float32(80)
	losslessCompression := 6

	err := processor.ProcessImages(
		config.Path,
		config.OutputPath,
		quality,
		config.Lossless,
		losslessCompression,
		config.Width,
		config.Height,
	)

	if err != nil {
		fmt.Printf("Error processing images: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Image processing completed successfully!")
}
