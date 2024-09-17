from scripts.models import Release

NIGHTLY_RELEASE: Release = {
    "version": "nightly",
    "name": "nightly",
    "platforms": {
        # "PlatformWindows": {
        #     "ArchitectureX86": {
        #         "downloadUrl": "https://artifacts.duckdb.org/latest/duckdb-binaries-windows.zip",
        #     },
        # },
        # "PlatformLinux": {
        #     "ArchitectureX86": {
        #         "downloadUrl": "https://artifacts.duckdb.org/latest/duckdb-binaries-linux.zip",
        #     },
        #     "ArchitectureArm64": {
        #         "downloadUrl": "https://artifacts.duckdb.org/latest/duckdb-binaries-linux-aarch64.zip",
        #     },
        # },
        # "PlatformMac": {
        #     "ArchitectureUniversal": {
        #         "downloadUrl": "https://artifacts.duckdb.org/latest/duckdb-binaries-osx.zip",
        #     },
        # },
    },
}
