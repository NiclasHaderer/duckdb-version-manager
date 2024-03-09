from typing import TypedDict, Literal

VersionList = dict[str, str]


class ArchitectureReleaseInformation(TypedDict):
    downloadUrl: str


ArchitectureType = Literal["ArchitectureX86", "ArchitectureArm64", "ArchitectureUniversal"]


PlatformType = Literal["PlatformWindows", "PlatformMac", "PlatformLinux"]


Architectures = dict[ArchitectureType, ArchitectureReleaseInformation]

Platforms = dict[PlatformType, Architectures]


class Release(TypedDict):
    version: str
    name: str
    platforms: Platforms


class AssetInformation(TypedDict):
    platform: PlatformType
    architecture: ArchitectureType
