# DuckDB Version Manager

### Installation

For **MacOS** and **Linux**, run the following command in your terminal:

```bash
curl https://raw.githubusercontent.com/NiclasHaderer/duckdb-version-manager/main/install.sh -s | /bin/bash
```

For **Windows**, download the latest binary from the [release page](https://github.com/NiclasHaderer/duckdb-version-manager/releases) and save it as **duckman.exe** in *$HOME/.local/bin*.
Then add *$HOME/.local/bin* to your PATH.

### Usage

```
duckman --help

A version manager for DuckDB

Usage:
  duckman [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  default     Set a version of DuckDB as default one to use.
  help        Help about any command
  install     Install a specific version of DuckDB
  list        List available DuckDB versions. Use 'local' to list local versions and 'remote' to list remote versions.
  run         Execute a specific version of DuckDB
  uninstall   Uninstall a version of DuckDB
  update-self  Updates duckman to the latest version

Flags:
  -h, --help   help for duckman

Use "duckman [command] --help" for more information about a command.
```

### Building from source

1. Install golang (https://go.dev/doc/install) if you haven't already
2. Run the *compile.sh* script to build binaries for all platforms. The binaries will be placed in the *bin* directory
   after the compilation is done.