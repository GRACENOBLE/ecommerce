#!/bin/bash
set -e

# Define target architecture for Linux
targets=("linux/amd64")

# Loop through each target
for target in "${targets[@]}"; do
    # Split the target into OS and architecture
    IFS="/" read -r goos goarch <<< "$target"
    
    # Define output binary name (adjust path and name as needed)
    output="bin/ecommerceBackend/ecommerceBackend-${goos}-${goarch}"
    
    echo "Building for $goos/$goarch..."
    
    # Set environment variables and build
    GOOS=$goos GOARCH=$goarch go build -tags netgo -ldflags '-s -w' -o "$output" ./cmd/ecommerceBackend

    chmod +x "$output"
    
done

echo "Linux amd64 build completed!"