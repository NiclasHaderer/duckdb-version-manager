import json
import os
import shutil
from collections import defaultdict
from datetime import datetime, timezone
from pathlib import Path
from typing import Literal

from github import Github
from github.Artifact import Artifact
from github.Auth import Token as GithubToken
from github.GitRelease import GitRelease
from github.PaginatedList import PaginatedList
from github.WorkflowRun import WorkflowRun
from packaging import version

from scripts.models import (
    VersionList,
    Release,
    AssetInformation,
    ArchitectureType,
    PlatformType,
)

VERSIONS_BASE_FOLDER = Path(__file__).parent.parent / "versions"
VERSIONS_INDIVIDUAL_FOLDER = VERSIONS_BASE_FOLDER / "tags"
shutil.rmtree(VERSIONS_BASE_FOLDER)
os.makedirs(VERSIONS_INDIVIDUAL_FOLDER, exist_ok=True)

auth_token = os.environ.get("GITHUB_AUTH_TOKEN")


def get_github() -> Github:
    if auth_token:
        return Github(auth=GithubToken(auth_token))
    else:
        return Github()


def get_all_releases_from(repo_name: str) -> PaginatedList[GitRelease]:
    g = get_github()
    repo = g.get_repo(repo_name)
    releases = repo.get_releases()
    return releases


def get_asset_type_from_name(asset_name: str, mode: Literal["duckman", "duck-db"]) -> AssetInformation | None:
    if mode == "duck-db" and (not asset_name.endswith(".zip") or "duckdb_cli" not in asset_name):
        return None

    if mode == "duckman" and asset_name.endswith(".zip"):
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


def save_duckdb_releases(nightly_release: Release):
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
        f.write(json.dumps(nightly_release, indent=2))

    version_list["nightly"] = "tags/nightly.json"

    with open(VERSIONS_BASE_FOLDER / "versions.json", "w") as f:
        f.write(json.dumps(version_list, indent=2))


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
        asset_info = get_asset_type_from_name(asset.name, "duckman")
        if asset_info is None:
            continue
        serializable_release["platforms"][asset_info["platform"]][asset_info["architecture"]] = {
            "downloadUrl": asset.browser_download_url
        }

    with open(VERSIONS_BASE_FOLDER / "latest-vm.json", "w") as f:
        f.write(json.dumps(serializable_release, indent=2))


def get_plain_url(run: WorkflowRun, artifact: Artifact) -> str:
    # https://github.com/duckdb/duckdb/actions/runs/10730276697/artifacts/1899513062
    url = "https://github.com/duckdb/duckdb/actions/runs"
    url += f"/{run.id}/artifacts/{artifact.id}"
    return url


def get_nightly_releases() -> Release:
    g = get_github()
    repo = g.get_repo("duckdb/duckdb")
    actions = repo.get_workflow_runs(branch=repo.get_branch("main"), exclude_pull_requests=True, status="success")

    windows: WorkflowRun | None = None
    linux: WorkflowRun | None = None
    osx: WorkflowRun | None = None
    idx = 0

    current_date = datetime.now(tz=timezone.utc)
    for flow in actions:
        idx += 1
        if (current_date - flow.created_at).days > 2 and idx > 500:
            break

        if windows and linux and osx:
            break

        if "nightly" not in flow.display_title:
            continue

        if "Windows" in flow.name and not windows:
            windows = flow
        elif "LinuxRelease" in flow.name and not linux:
            linux = flow
        elif "OSX" in flow.name and not osx:
            osx = flow

    release = {"version": "nightly", "name": "nightly", "platforms": {}}

    if windows:
        artifacts = windows.get_artifacts()
        release["platforms"]["PlatformWindows"] = {}
        for artifact in artifacts:
            if "amd64" in artifact.name:
                release["platforms"]["PlatformWindows"]["ArchitectureX86"] = {
                    "downloadUrl": get_plain_url(windows, artifact)
                }
            elif "arm64" in artifact.name:
                release["platforms"]["PlatformWindows"]["ArchitectureArm64"] = {
                    "downloadUrl": get_plain_url(windows, artifact)
                }

    if osx:
        artifacts = osx.get_artifacts()
        release["platforms"]["PlatformMac"] = {}
        for artifact in artifacts:
            if "binaries" in artifact.name:
                release["platforms"]["PlatformMac"]["ArchitectureUniversal"] = {
                    "downloadUrl": get_plain_url(osx, artifact)
                }

    if linux:
        artifacts = linux.get_artifacts()
        release["platforms"]["PlatformLinux"] = {}
        for artifact in artifacts:
            if "aarch64" in artifact.name and "binaries" in artifact.name:
                release["platforms"]["PlatformLinux"]["ArchitectureArm64"] = {
                    "downloadUrl": get_plain_url(linux, artifact)
                }
            elif "binaries" in artifact.name:
                release["platforms"]["PlatformLinux"]["ArchitectureX86"] = {
                    "downloadUrl": get_plain_url(linux, artifact)
                }

    return release


def main():
    save_latest_duck_vm_release()
    nightly_release = get_nightly_releases()
    save_duckdb_releases(nightly_release)


if __name__ == "__main__":
    main()
