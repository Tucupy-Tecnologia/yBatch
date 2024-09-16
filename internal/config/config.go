package config

import (
	"flag"
	"fmt"
	"os"
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

	// Check if there are any arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a path to the directory containing images as the first argument.")
		os.Exit(1)
	}

	// Set the path from the first argument
	config.Path = os.Args[1]

	// Create a new FlagSet
	fs := flag.NewFlagSet("yBatch", flag.ExitOnError)

	// Define flags
	fs.StringVar(&config.OutputPath, "output", "", "Output directory for the processed images")
	fs.IntVar(&config.Width, "w", 0, "Width of the output images")
	fs.IntVar(&config.Height, "h", 0, "Height of the output images")
	fs.StringVar(&config.Format, "format", "webp", "Output format for the images (e.g., webp, jpg, png)")
	fs.StringVar(&config.AspectRatio, "ar", "", "Aspect ratio for the output images (e.g., 16:9)")
	fs.BoolVar(&config.Lossless, "lossless", false, "Enable lossless WebP encoding")

	// Parse the remaining flags
	fs.Parse(os.Args[2:])

	return config
}
