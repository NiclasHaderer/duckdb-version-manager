from scripts.models import Release

NIGHTLY_RELEASE: Release = {
    "version": "nightly",
    "name": "nightly",
    "platforms": {
        "PlatformWindows": {
            "ArchitectureX86": {
                "downloadUrl": "https://artifacts.duckdb.org/latest/duckdb-binaries-windows.zip",
            },
        },
        "PlatformLinux": {
            "ArchitectureX86": {
                "downloadUrl": "https://artifacts.duckdb.org/latest/duckdb-binaries-linux-amd64.zip",
            },
            "ArchitectureArm64": {
                "downloadUrl": "https://artifacts.duckdb.org/latest/duckdb-binaries-linux-arm64.zip",
            },
        },
        "PlatformMac": {
            "ArchitectureUniversal": {
                "downloadUrl": "https://artifacts.duckdb.org/latest/duckdb-binaries-osx.zip",
            },
        },
    },
}
