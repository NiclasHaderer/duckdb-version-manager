import json
import os
import shutil
from collections import defaultdict
from pathlib import Path
from typing import Literal

import github.Auth
import httpx
from github import Github
from github.GitRelease import GitRelease
from github.PaginatedList import PaginatedList
from packaging import version

from scripts.models import (
    VersionList,
    Release,
    AssetInformation,
    ArchitectureType,
    PlatformType,
)
from scripts.nightly import NIGHTLY_RELEASE

VERSIONS_BASE_FOLDER = Path(__file__).parent.parent / "versions"
VERSIONS_INDIVIDUAL_FOLDER = VERSIONS_BASE_FOLDER / "tags"
shutil.rmtree(VERSIONS_BASE_FOLDER)
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


def get_asset_type_from_name(asset_name: str, mode: Literal["duck-vm", "duck-db"]) -> AssetInformation | None:
    if mode == "duck-db" and (not asset_name.endswith(".zip") or "duckdb_cli" not in asset_name):
        return None

    if mode == "duck-vm" and asset_name.endswith(".zip"):
        return None

    architecture: ArchitectureType
    if "amd64" in asset_name:
        architecture = "ArchitectureX86"
    elif "aarch64" in asset_name or "arm64" in asset_name:
        architecture = "ArchitectureArm64"
    elif "universal" in asset_name:
        architecture = "ArchitectureUniversal"
    else:
        return None

    platform: PlatformType
    if "windows" in asset_name:
        platform = "PlatformWindows"
    elif "osx" in asset_name or "darwin" in asset_name:
        platform = "PlatformMac"
    elif "linux" in asset_name:
        platform = "PlatformLinux"
    else:
        return None

    return {"platform": platform, "architecture": architecture}


def save_duckdb_releases():
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
            asset_info = get_asset_type_from_name(asset.name, "duck-db")
            if asset_info is None:
                continue
            serializable_release["platforms"][asset_info["platform"]][asset_info["architecture"]] = {
                "downloadUrl": asset.browser_download_url
            }

        if not serializable_release["platforms"]:
            continue

        with open(VERSIONS_INDIVIDUAL_FOLDER / f"{tag_name}.json", "w") as f:
            f.write(json.dumps(serializable_release, indent=2))

    with open(VERSIONS_INDIVIDUAL_FOLDER / f"nightly.json", "w") as f:
        f.write(json.dumps(NIGHTLY_RELEASE, indent=2))

    version_list["nightly"] = "tags/nightly.json"

    with open(VERSIONS_BASE_FOLDER / "versions.json", "w") as f:
        f.write(json.dumps(version_list, indent=2))


def validate_nightly_urls(nightly: Release):
    for platform, architectures in nightly["platforms"].items():
        for architecture, info in architectures.items():
            if httpx.get(info["downloadUrl"], follow_redirects=True).status_code != 200:
                raise ValueError(f"URL {info['downloadUrl']} is not valid")


def save_latest_duck_vm_release():
    repo_name = "niclashaderer/duckdb-version-manager"
    releases = get_all_releases_from(repo_name)
    latest_release = releases[0]
    for release in releases:
        if version.parse(release.tag_name) > version.parse(latest_release.tag_name):
            latest_release = release

    serializable_release: Release = {
        "version": latest_release.tag_name,
        "name": latest_release.title,
        "platforms": defaultdict(dict),
    }
    for asset in latest_release.get_assets():
        asset_info = get_asset_type_from_name(asset.name, "duck-vm")
        if asset_info is None:
            continue
        serializable_release["platforms"][asset_info["platform"]][asset_info["architecture"]] = {
            "downloadUrl": asset.browser_download_url
        }

    with open(VERSIONS_BASE_FOLDER / "latest-vm.json", "w") as f:
        f.write(json.dumps(serializable_release, indent=2))


def main():
    validate_nightly_urls(NIGHTLY_RELEASE)
    save_latest_duck_vm_release()
    save_duckdb_releases()


if __name__ == "__main__":
    main()
