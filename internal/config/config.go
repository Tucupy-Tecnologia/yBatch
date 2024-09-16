package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Path        string
	OutputPath  string
	Width       int
	Height      int
	Format      string
	AspectRatio string
	Lossless    bool
}

func ParseFlags() Config {
	var config Config

	// Define flags
	flag.StringVar(&config.Path, "path", "", "Path to the directory containing images")
	flag.StringVar(&config.AspectRatio, "ar", "", "Aspect ratio for the output images (e.g., 16:9)")
	flag.StringVar(&config.Format, "format", "webp", "Output format for the images (e.g., webp, jpg, png)")
	flag.BoolVar(&config.Lossless, "lossless", false, "Enable lossless WebP encoding")
	flag.IntVar(&config.Width, "width", 0, "Width of the output images")
	flag.IntVar(&config.Height, "height", 0, "Height of the output images")
	flag.StringVar(&config.OutputPath, "output", "", "Output directory for the processed images")

	// Parse flags
	flag.Parse()

	// Check if the first argument is provided and not a flag
	args := flag.Args()
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		config.Path = args[0]
	}

	// If path is still empty, prompt the user
	if config.Path == "" {
		fmt.Println("Please provide a valid path to a folder containing the files to be processed")
		fmt.Println("You can do this by:")
		fmt.Println("1. Passing it as the first argument: yBatch /path/to/folder")
		fmt.Println("2. Using the --path flag: yBatch --path /path/to/folder")
		os.Exit(1)
	}

	return config
}
