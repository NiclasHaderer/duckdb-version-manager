VERSION="--dev--"
OUTPUT_NAME="$HOME/.local/bin/duckman"
go build -ldflags "-X 'duckdb-version-manager/cmd.version=$VERSION'" -o "$OUTPUT_NAME"