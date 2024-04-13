import os
import re
import subprocess
import sys
from pathlib import Path

main_file = str(Path(__file__).parent.parent / "main.go")


def run_process(*args: str) -> str:
    args = ["go", "run", "-ldflags", "-X 'duckdb-version-manager/config.Version=100.0.0'", main_file, *args]
    result = subprocess.run(
        args,
        capture_output=True,
        text=True,
    )

    text = result.stdout + result.stderr
    if result.returncode != 0:
        raise ValueError(text)
    return text


def install_version(version: str) -> str:
    return run_process("install", version)


def list_local_versions() -> list[str]:
    out = run_process("list", "local")
    regex = re.compile(r"^.*?((?:\d|\.|v){3,}|nightly)", re.M)
    installed_versions = re.findall(regex, out)
    return installed_versions


def list_remote_versions() -> list[str]:
    out = run_process("list", "remote")
    regex = re.compile(r"^.*?((?:\d|\.|v){3,}|nightly)", re.M)
    remote_versions = re.findall(regex, out)
    return remote_versions


def set_version_as_default(version: str) -> str:
    return run_process("default", version)


def run_version(version: str, *args: str) -> str:
    return run_process("run", version, *args)


def uninstall_version(version: str) -> str:
    return run_process("uninstall", version)


def run_default(*args: str) -> str:
    home_dir = Path.home()
    install_dir = home_dir / ".local" / "bin"
    binary = "duckdb"
    if sys.platform == "win32":
        binary += ".exe"

    binary = install_dir / binary
    result = subprocess.run([str(binary), *args], capture_output=True, text=True)
    return result.stdout
