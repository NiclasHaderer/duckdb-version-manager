VERSION="100.0.0"
OUTPUT_NAME="$HOME/.local/bin/duckman"
go build -ldflags "-X 'duckdb-version-manager/config.Version=$VERSION'" -o "$OUTPUT_NAME"