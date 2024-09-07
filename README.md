# Manga Hub

## Overview

Manga Hub is a Go-based project for managing and handling manga data. This README will guide you through the setup process, including building the project, running it in development and production modes, and cleaning up build artifacts.

## Project Structure

- `cmd/` - Contains the main application entry point.
- `handlers/` - Contains handlers for HTTP requests.
- `config/` - Contains configuration files.
- `build.sh` - Script for building and managing the project.
- `go.mod` - Go module file for dependency management.
- `go.sum` - Go module checksum file for dependency management.

## Setup

### Prerequisites

1. **Go**: Ensure you have Go installed. You can download and install Go from the [official Go website](https://golang.org/dl/).
2. **Air**: For hot reloading in development, install [Air](https://github.com/cosmtrek/air).

### Cloning the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/sahiljaggarwal/go_manga_hub.git
cd repository
```

### Installing Dependencies

Install the necessary Go dependencies:

```bash
go mod tidy
```

### Build Script

The `build.sh` script automates various tasks:

- **`build_dir`**: Creates the build directory.
- **`build`**: Builds the project and outputs the binary to the `build` directory.
- **`dev`**: Runs the project in development mode using Air for hot reloading.
- **`prod`**: Builds the project and then runs the binary in production mode.
- **`clean`**: Cleans up the build directory.

### Running the Build Script

You can use the `build.sh` script to perform different tasks. Run the script with one of the following arguments:

- **Create Build Directory**:

  ```bash
  ./build.sh build_dir
  ```

- **Build the Project**:

  ```bash
  ./build.sh build
  ```

- **Run in Development Mode**:

  ```bash
  ./build.sh dev
  ```

- **Run in Production Mode**:

  ```bash
  ./build.sh prod
  ```

- **Clean the Build Directory**:

  ```bash
  ./build.sh clean
  ```

### Configuration

The project configuration is handled through environment variables. Create a `.env` file in the root directory with the following variables:

```env
PORT=8080
```

### Running the Project

- **Development Mode**: Use `./build.sh dev` to start the project with hot reloading.
- **Production Mode**: Use `./build.sh prod` to build the project and run the binary.

### Logging

The project configuration includes logging settings. Modify the `air` configuration file to suit your logging preferences:

```toml
[tmp]
bin = "./build/manga-hub.exe"
cmd = "go build -o ./build/manga-hub.exe ./cmd"

[log]
time = true
color = true
```

## Handlers

The `handlers` package includes functionality for processing HTTP requests and managing images. It utilizes `goquery` for HTML parsing, `gin` for web framework functionality, and `gopdf` for PDF generation.

## Contributing

Contributions are welcome! Please fork the repository, make your changes, and create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

