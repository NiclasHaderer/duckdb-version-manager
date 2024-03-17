import re
import subprocess
from pathlib import Path

main_file = str(Path(__file__).parent.parent / "main.go")


def run_process(*args: str) -> str:
    args = ["go", "run", "-ldflags", "-X 'duckdb-version-manager/cmd.version=dev'", main_file, *args]
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
    regex = r"  ((?:\w|\.|[09])+)"
    return re.findall(regex, out)


def list_remote_versions() -> list[str]:
    out = run_process("list", "remote")
    regex = r"  ((?:\w|\.|[09])+)"
    return re.findall(regex, out)


def set_version_as_default(version: str) -> str:
    return run_process("default", version)


def run_version(version: str, *args: str) -> str:
    return run_process("run", version, *args)


def uninstall_version(version: str) -> str:
    return run_process("uninstall", version)


def run_default(*args: str) -> str:
    result = subprocess.run(["duckdb", *args], capture_output=True, text=True)
    return result.stdout
