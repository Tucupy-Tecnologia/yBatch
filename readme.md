# yBatch

yBatch is a command-line tool for batch processing images. It allows users to resize images and convert them to WebP format efficiently.

## Features

- Batch process images in a directory
- Resize images to specified dimensions
- Convert images to WebP format
- Support for lossless WebP encoding
- Configurable output directory

## Installation

**Note**: yBatch requires Go 1.22 or later.
 
If you don't have Go installed, you can download it from the official website: https://go.dev/dl/

```homebew installation coming soon```

To install yBatch, you need to have Go installed on your system. Then, you can clone the repository and build the project:

```bash
git clone https://github.com/Tucupy-Tecnologia/yBatch.git
cd yBatch
go build -o yBatch cmd/main.go
```

This will create an executable named `yBatch` in your current directory.

## Usage

The basic syntax for using yBatch is:

```bash
./yBatch <path_to_image_directory> [flags]
```

### Flags

- `-w`: Width of the output images (int)
- `-h`: Height of the output images (int)
- `--output`: Output directory for the processed images (string)
- `--format`: Output format for the images (default "webp")
- `--ar`: Aspect ratio for the output images (e.g., "16:9")
- `--lossless`: Enable lossless WebP encoding (bool)

### Examples

1. Resize images to 400x400 pixels:
   ```bash
   ./yBatch /path/to/images -w 400 -h 400
   ```

2. Convert images to WebP format without resizing:
   ```bash
   ./yBatch /path/to/images
   ```

3. Use lossless encoding:
   ```bash
   ./yBatch /path/to/images --lossless
   ```

4. Specify an output directory:
   ```bash
   ./yBatch /path/to/images --output /path/to/output
   ```

## How It Works

1. The program takes the first argument as the path to the directory containing images.
2. It then parses the remaining flags to determine the processing options.
3. Each image in the specified directory is processed concurrently:
   - If width and height are specified, the image is resized.
   - The image is then converted to WebP format (either lossless or lossy, depending on the `--lossless` flag).
4. Processed images are saved to the output directory (if specified) or to a new directory named "yBatch-output-[original_directory_name]" in the parent directory of the input.

## Dependencies

yBatch uses the following external libraries:

- github.com/kolesa-team/go-webp: For WebP encoding
- github.com/nfnt/resize: For image resizing

## Contributing

Contributions to yBatch are welcome! Please feel free to submit a Pull Request.

## License

yBatch is licensed under the MIT License. See the LICENSE file for more information. 

## Support

If you encounter any problems or have any questions, please open an issue on the GitHub repository.
