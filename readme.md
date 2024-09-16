# yBatch

yBatch is a command-line tool for batch processing images. It allows users to resize images and convert them to WebP format efficiently.

## Features

- Batch process images in a directory
- Resize images to specified dimensions
- Convert images to WebP format
- Support for lossless WebP encoding
- Configurable output directory

## Installation

### Prerequisites

- Go (version 1.22.0 or later)
- Git

### Local Installation

To install yBatch locally:

1. Clone the repository:
   ```bash
   git clone https://github.com/Tucupy-Tecnologia/yBatch.git
   ```

2. Navigate to the project directory:
   ```bash
   cd yBatch
   ```

3. Build the project:
   ```bash
   go build -o ybatch cmd/main.go
   ```

This will create an executable named `ybatch` in your current directory.

### Global Installation

To use `ybatch` from any directory, you can install it globally. The process varies depending on your operating system.

#### For macOS (including Warp terminal users):

1. Build the program (if you haven't already):
   ```bash
   go build -o ybatch cmd/main.go
   ```

2. Move the executable to a directory in your PATH:
   ```bash
   sudo mv ybatch /usr/local/bin/
   ```

3. Set the correct permissions:
   ```bash
   sudo chmod +x /usr/local/bin/ybatch
   ```

4. Update your shell configuration. For Zsh (default in macOS and Warp):
   ```bash
   echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.zshrc
   source ~/.zshrc
   ```

5. For Warp terminal users, you might need to update the Warp configuration:
   ```bash
   echo 'env:' >> ~/.warp/launch_configurations.yml
   echo '  PATH: /usr/local/bin:$PATH' >> ~/.warp/launch_configurations.yml
   ```

6. Restart your terminal or Warp application.

7. Verify the installation:
   ```bash
   which ybatch
   ```
   This should output `/usr/local/bin/ybatch`.

#### For Linux:

1. Build the program:
   ```bash
   go build -o ybatch cmd/main.go
   ```

2. Move the executable to a directory in your PATH:
   ```bash
   sudo mv ybatch /usr/local/bin/
   ```

3. Set the correct permissions:
   ```bash
   sudo chmod +x /usr/local/bin/ybatch
   ```

4. Update your shell configuration (for bash):
   ```bash
   echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.bashrc
   source ~/.bashrc
   ```

#### For Windows:

1. Build the program:
   ```bash
   go build -o ybatch.exe cmd/main.go
   ```

2. Move the `ybatch.exe` file to `C:\Windows\` or another directory in your PATH.

3. You may need to restart your command prompt or PowerShell for the changes to take effect.

## Usage

The basic syntax for using yBatch is:

```bash
ybatch <path_to_image_directory> [flags]
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
   ybatch /path/to/images -w 400 -h 400
   ```

2. Convert images to WebP format without resizing:
   ```bash
   ybatch /path/to/images
   ```

3. Use lossless encoding:
   ```bash
   ybatch /path/to/images --lossless
   ```

4. Specify an output directory:
   ```bash
   ybatch /path/to/images --output /path/to/output
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

yBatch is licensed under the MIT License.

## Support

If you encounter any problems or have any questions, please open an issue on the GitHub repository.

## Troubleshooting

If you're having trouble running `ybatch` after installation:

1. Ensure the executable is in a directory listed in your PATH:
   ```bash
   echo $PATH
   ```

2. Verify the installation location:
   ```bash
   which ybatch
   ```

3. Check if the executable has the correct permissions:
   ```bash
   ls -l $(which ybatch)
   ```

4. If you've made changes to your shell configuration or Warp configuration, make sure to restart your terminal or source the updated configuration file.

If problems persist, please open an issue on the GitHub repository with details about your system and the error you're encountering.
