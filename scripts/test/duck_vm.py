import re
from pathlib import Path

from pwnlib.tubes import process

main_file = str(Path(__file__).parent.parent.parent / "main.go")
print(main_file)


def run_process(*args: str) -> str:
    p = process.process(["go", "run", main_file, *args])
    stdout = p.recvall().decode("utf-8")
    if "Duck-VM encounted a fatal error" in stdout:
        raise ValueError(stdout)
    return stdout


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
    p = process.process(["duckdb", *args])
    return p.recvall().decode("utf-8")
