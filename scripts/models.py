from enum import Enum, auto
from typing import TypedDict

VersionList = dict[str, str]


class ArchitectureReleaseInformation(TypedDict):
    downloadUrl: str


Architectures = dict[str, ArchitectureReleaseInformation]

Platforms = dict[str, Architectures]


class Release(TypedDict):
    version: str
    name: str
    platforms: Platforms


class ArchitectureType(Enum):
    ArchitectureX86 = auto()
    ArchitectureArm64 = auto()
    ArchitectureUniversal = auto()


class PlatformType(Enum):
    PlatformWindows = auto()
    PlatformMac = auto()
    PlatformLinux = auto()


class AssetInformation(TypedDict):
    platform: PlatformType
    architecture: ArchitectureType
