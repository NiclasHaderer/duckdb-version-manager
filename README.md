# DuckDB Version Manager

[![Tests](https://github.com/NiclasHaderer/duckdb-version-manager/actions/workflows/test.yml/badge.svg)](https://github.com/NiclasHaderer/duckdb-version-manager/actions/workflows/test.yml)

### Installation

> **Note:** Before proceeding with the installation, it's advisable to uninstall or remove any existing DuckDB CLI versions that you may have installed manually or via brew.

For **macOS** and **Linux**, run the following command in your terminal:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/NiclasHaderer/duckdb-version-manager/main/install.sh)"
```

For **Windows**, run the following command in your powershell terminal:

```powershell
Invoke-Expression (Invoke-WebRequest -UseBasicParsing -Uri "https://raw.githubusercontent.com/NiclasHaderer/duckdb-version-manager/main/install.ps1").Content
```

### Usage

```bash
# Install a specific version of DuckDB (this does not set the version as default)
duckman install 0.10.2

# Set a version of DuckDB as default one to use -> now running duckdb will run this version
duckman default 0.10.2

# Run a version of DuckDB
duckman run nightly

# List available DuckDB versions
duckman list remote
```

Generally, installing a version before running it or setting it as default is not necessary.
If you want to run a version that is not installed, duckman will automatically download and install it for you.

### Reference


```
A version manager for DuckDB

Usage:
  duckman [command]

Available Commands:
  completion     Generate the autocompletion script for the specified shell
  default        Set a version of DuckDB as default one to use.
  help           Help about any command
  install        Install a specific version of DuckDB
  list           List available DuckDB versions. Use 'local' to list local versions and 'remote' to list remote versions.
  run            Execute a specific version of DuckDB
  uninstall      Uninstall a version of DuckDB
  uninstall-self Removes duckman and all config files
  update-self    Updates duckman to the latest version

Flags:
  -h, --help      help for duckman
  -v, --version   version for duckman

Use "duckman [command] --help" for more information about a command.
```

### Building from source

1. Install golang (https://go.dev/doc/install) if you haven't already
2. Run the *compile.sh* script to build binaries for all platforms. The binaries will be placed in the *bin* directory
   after the compilation is done.
