#!/bin/bash
set -e

# Define target architectures in the format "GOOS/GOARCH"
targets=("linux/amd64" "darwin/amd64" "windows/amd64")

# Loop through each target
for target in "${targets[@]}"; do
    # Split the target into OS and architecture
    IFS="/" read -r goos goarch <<< "$target"
    
    # Define output binary name (adjust path and name as needed)
    output="bin/ecommerceBackend/ecommerceBackend-${goos}-${goarch}"
    
    echo "Building for $goos/$goarch..."
    
    # Set environment variables and build
    GOOS=$goos GOARCH=$goarch go build -o "$output" ./cmd/ecommerceBackend
done

echo "Builds completed!"
