import json
import os
from collections import defaultdict
from pathlib import Path

import github.Auth
from github import Github
from github.GitRelease import GitRelease
from github.PaginatedList import PaginatedList

from scripts.models import (
    VersionList,
    Release,
    AssetInformation,
    ArchitectureType,
    PlatformType,
)

VERSIONS_BASE_FOLDER = Path(__file__).parent.parent / "versions"
VERSIONS_INDIVIDUAL_FOLDER = VERSIONS_BASE_FOLDER / "tags"
os.makedirs(VERSIONS_INDIVIDUAL_FOLDER, exist_ok=True)

auth_token = os.environ.get("GITHUB_AUTH_TOKEN")


def get_all_releases_from(repo_name: str) -> PaginatedList[GitRelease]:
    g: Github
    if auth_token:
        g = Github(auth=github.Auth.Token(auth_token))
    else:
        g = Github()
    repo = g.get_repo(repo_name)
    releases = repo.get_releases()
    return releases


def get_asset_type_from_name(asset_name: str) -> AssetInformation | None:
    if not asset_name.endswith(".zip") or "duckdb_cli" not in asset_name:
        return None
    architecture: ArchitectureType
    if "amd64" in asset_name:
        architecture = ArchitectureType.ArchitectureX86
    elif "aarch64" in asset_name:
        architecture = ArchitectureType.ArchitectureArm64
    elif "universal" in asset_name:
        architecture = ArchitectureType.ArchitectureUniversal
    else:
        return None

    platform: PlatformType
    if "windows" in asset_name:
        platform = PlatformType.PlatformWindows
    elif "osx" in asset_name:
        platform = PlatformType.PlatformMac
    elif "linux" in asset_name:
        platform = PlatformType.PlatformLinux
    else:
        return None

    return {"platform": platform, "architecture": architecture}


def main():
    repo_name = "duckdb/duckdb"
    releases = get_all_releases_from(repo_name)

    version_list: VersionList = {}
    for release in releases:
        tag_name = release.tag_name
        version_list[tag_name] = f"tags/{tag_name}.json"

        serializable_release: Release = {
            "version": tag_name,
            "name": release.title,
            "platforms": defaultdict(dict),
        }
        for asset in release.get_assets():
            asset_info = get_asset_type_from_name(asset.name)
            if asset_info is None:
                continue
            serializable_release["platforms"][asset_info["platform"].name][asset_info["architecture"].name] = {
                "downloadUrl": asset.browser_download_url
            }

        with open(VERSIONS_INDIVIDUAL_FOLDER / f"{tag_name}.json", "w") as f:
            f.write(json.dumps(serializable_release, indent=2))

    with open(VERSIONS_BASE_FOLDER / "versions.json", "w") as f:
        f.write(json.dumps(version_list, indent=2))


if __name__ == "__main__":
    main()
