#!/bin/bash

# Get version from git
VERSION=$1

# Check if version is set
if [ -z "$VERSION" ]; then
	echo "Version not set"
	exit 1
fi

# Variables
BINARY_NAME="duckman"
OUTPUT_DIR="bin"
GOOS_ARRAY=("darwin" "darwin" "linux" "linux" "windows" "windows")
GOARCH_ARRAY=("amd64" "arm64" "amd64" "arm64" "amd64" "arm64")

# Create output directory
rm -r $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

# Loop through architectures
# shellcheck disable=SC2068
for i in ${!GOOS_ARRAY[@]}; do
	GOOS=${GOOS_ARRAY[$i]}
	GOARCH=${GOARCH_ARRAY[$i]}

	OUTPUT_NAME="$OUTPUT_DIR/$BINARY_NAME-$GOOS-$GOARCH"

	# Cross-compile
	env GOOS="$GOOS" GOARCH="$GOARCH" go build -ldflags "-X 'duckdb-version-manager/cmd.version=$VERSION'" -o "$OUTPUT_NAME"

	# Check if cross-compilation was successful
	if [ $? -eq 0 ]; then
		echo "Successfully compiled $BINARY_NAME for $GOOS $GOARCH"
	else
		echo "Error compiling $BINARY_NAME for $GOOS $GOARCH"
		exit 1
	fi
done
