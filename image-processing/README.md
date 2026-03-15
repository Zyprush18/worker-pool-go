# Image Processing Worker Pool

A Go implementation of a worker pool pattern for high-performance batch image processing. This project demonstrates concurrent image operations using goroutines and channels.

## Features

- **Image Conversion**: Convert images to JPEG format
- **Image Resizing**: Resize images to 300x400 pixels
- **Image Compression**: Compress images with quality 80 and max size 2048
- **Concurrent Processing**: Uses worker pool pattern for efficient parallel processing
- **Channel-based Communication**: Leverages Go channels for goroutine synchronization

## Requirements

- Go 1.25.4 or later
- ImageMagick (required by bimg library)

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Zyprush18/worker-pool-go
   cd image-processing
   ```

2. **Install ImageMagick** (required for bimg)
   
   On Ubuntu/Debian:
   ```bash
   sudo apt-get install libmagickwand-dev
   ```
   
   On macOS:
   ```bash
   brew install imagemagick
   ```

3. **Install Go dependencies**
   ```bash
   go mod download
   ```

## Setup

Before running the project, you need to create the required directories:

```bash
# Create input image directory
mkdir -p image

# Create output directories for processed images
mkdir -p image-compress
mkdir -p image-convert
mkdir -p image-resize
```

## Usage

1. **Add images to process**
   
   Place your images in the `image/` directory. Supported formats include JPEG, PNG, WebP, and AVIF.

2. **Run the image processor**

   The program accepts one argument to specify the operation:

   ```bash
   # Convert images to JPEG
   go run main.go worker.go convert

   # Resize images to 300x400 pixels
   go run main.go worker.go resize

   # Compress images (quality: 80, max size: 2048)
   go run main.go worker.go compress
   ```

3. **Find processed images**

   - **Converted images**: `image-convert/`
   - **Resized images**: `image-resize/`
   - **Compressed images**: `image-compress/`

## Directory Structure

```
.
├── image/                 # Input images directory
├── image-compress/        # Compressed output
├── image-convert/         # Converted output
├── image-resize/          # Resized output
├── main.go               # Entry point
├── worker.go             # Worker pool implementation
├── go.mod                # Go module file
└── README.md             # This file
```

## How It Works

1. **Main Function**: Reads images from the `image/` directory
2. **Channels**: Creates channels for image data and filenames
3. **Worker Pool**: Spawns 3 concurrent workers to process images
4. **Processing**: Each worker applies the specified transformation (convert/resize/compress)
5. **Output**: Processed images are saved to the appropriate output directory

## Dependencies

- **bimg** (github.com/h2non/bimg): Go bindings for ImageMagick

## Building

```bash
# Build the executable
go build -o image-processor main.go worker.go

# Run the built executable
./image-processor convert
```

## Example

```bash
# Process images in the image/ directory
go run main.go worker.go convert
# Check image-convert/ for the converted JPEG files

go run main.go worker.go resize
# Check image-resize/ for the resized images

go run main.go worker.go compress
# Check image-compress/ for the compressed images
```

## Author

Zyprush18
