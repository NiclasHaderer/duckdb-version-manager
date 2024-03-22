#!/bin/bash

URL="https://raw.githubusercontent.com/NiclasHaderer/duckdb-version-manager/main/versions/latest-vm.json"

if ! command -v jq &> /dev/null; then
    echo "jq could not be found. Please install jq to proceed."
    exit 1
fi

DOWNLOAD_DIR="$HOME/.local/bin"
mkdir -p "${DOWNLOAD_DIR}"

JSON_CONTENT=$(curl -sL "${URL}")

OS="$(uname -s)"
ARCH="$(uname -m)"

case "${ARCH}" in
    x86_64) ARCH="ArchitectureX86" ;;
    arm64) ARCH="ArchitectureArm64" ;;
    *) echo "Unsupported architecture: ${ARCH}"; exit 2 ;;
esac

case "${OS}" in
    Linux) OS_KEY="PlatformLinux" ;;
    Darwin) OS_KEY="PlatformMac" ;;
    *) echo "Unsupported OS: ${OS}"; exit 3 ;;
esac

DOWNLOAD_URL=$(echo "${JSON_CONTENT}" | jq -r ".platforms.${OS_KEY}.${ARCH}.downloadUrl")

if [ "${DOWNLOAD_URL}" == "null" ] || [ -z "${DOWNLOAD_URL}" ]; then
    echo "Failed to find a valid download URL for your platform."
    exit 4
fi

echo "Downloading from ${DOWNLOAD_URL}..."
curl -sL "${DOWNLOAD_URL}" -o "${DOWNLOAD_DIR}/duckman"
echo "Download complete. duckman is now available in ${DOWNLOAD_DIR}/duckman"

chmod +x "${DOWNLOAD_DIR}/duckman"

if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
    echo ""
    echo "$HOME/.local/bin is not in PATH."
    echo "Add it to your PATH by adding the following line to your shell's configuration file:"
    echo "export PATH=\"\$HOME/.local/bin:\$PATH\""
fi

echo "To get autocomplete for duckman, add the following line to your shell's configuration file:"
echo "eval \"\$(duckman completion <zsh|bash|fish>)\""