#!/bin/bash

# Function to create the build directory
build_dir() {
  mkdir -p build
  echo "Build directory created."
}

# Function to build the project
build() {
  build_dir
  go build -o build/manga-hub cmd/main.go
  echo "Project built successfully."
}

# Function to run the project in development mode
dev() {
  air
}

# Function to run the project in production mode
prod() {
  build
  ./build/manga-hub
}

# Function to clean the build directory
clean() {
  rm -rf build
  echo "Build directory cleaned."
}

# Main logic to handle script arguments
case "$1" in
  build_dir)
    build_dir
    ;;
  build)
    build
    ;;
  dev)
    dev
    ;;
  prod)
    prod
    ;;
  clean)
    clean
    ;;
  *)
    echo "Usage: $0 {build_dir|build|dev|prod|clean}"
    exit 1
    ;;
esac
