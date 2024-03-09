#!/bin/bash

# Variables
BINARY_NAME="duck-vm"
OUTPUT_DIR="bin"
GOOS_ARRAY=("darwin" "darwin" "linux" "linux" "windows" "windows")
GOARCH_ARRAY=("amd64" "arm64" "amd64" "arm" "amd64" "arm")

# Create output directory
mkdir -p $OUTPUT_DIR

# Loop through architectures
# shellcheck disable=SC2068
for i in ${!GOOS_ARRAY[@]}; do
    GOOS=${GOOS_ARRAY[$i]}
    GOARCH=${GOARCH_ARRAY[$i]}

    OUTPUT_NAME="$OUTPUT_DIR/$BINARY_NAME-$GOOS-$GOARCH"

    # Cross-compile
    env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$OUTPUT_NAME"

    # Check if cross-compilation was successful
    if [ $? -eq 0 ]; then
        echo "Successfully compiled $BINARY_NAME for $GOOS $GOARCH"
    else
        echo "Error compiling $BINARY_NAME for $GOOS $GOARCH"
        exit 1
    fi
done