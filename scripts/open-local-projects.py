"""Build missing previews, keep them running, and open the four local projects."""
import argparse
import json
import os
from pathlib import Path
import shutil
import signal
import subprocess
import time
from urllib.request import urlopen

PORTFOLIO = Path(__file__).resolve().parents[1]
STATE = PORTFOLIO / ".local-state"


def healthy(url, endpoint):
    try:
        with urlopen(url + endpoint, timeout=2) as response:
            return response.status == 200 and json.load(response).get("status") == "ok"
    except (OSError, ValueError):
        return False


def owned_process(info):
    """Never signal an unrelated process after a restart or PID reuse."""
    try:
        process = Path("/proc") / str(info["pid"])
        return (process / "cwd").resolve() == Path(info["cwd"]) and (
            process / "cmdline"
        ).read_bytes().split(b"\0")[:-1] == [arg.encode() for arg in info["command"]]
    except OSError:
        return False


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--project-root", type=Path, default=PORTFOLIO.parent / "projects")
    parser.add_argument("--no-open", action="store_true", help="Start without browser tabs")
    parser.add_argument("--stop", action="store_true", help="Stop only previews started here")
    parser.add_argument("--rebuild", action="store_true", help="Rebuild and restart owned previews")
    args = parser.parse_args()
    STATE.mkdir(exist_ok=True)
    registry = STATE / "previews.json"
    processes = json.loads(registry.read_text()) if registry.exists() else {}
    if args.stop or args.rebuild:
        for name, info in processes.items():
            if owned_process(info):
                os.kill(info["pid"], signal.SIGTERM)
                print("Stopped " + name)
        processes = {}
        registry.write_text("{}\n")
        if args.stop:
            return
        time.sleep(.5)

    node = shutil.which("node")
    if not node:
        candidate = Path.home() / ".local/node/bin/node"
        if candidate.is_file():
            node = str(candidate)
    if not node:
        raise RuntimeError("Install Node 22.13+ and npm before starting the applications.")
    env = {**os.environ, "PATH": str(Path(node).parent) + os.pathsep + os.environ.get("PATH", "")}
    npm = shutil.which("npm", path=env["PATH"])
    if not npm:
        raise RuntimeError("npm is required to build the applications.")
    previews = [("portfolio", PORTFOLIO, 8085, "/healthz")]
    previews += [(name, args.project_root.resolve() / name, port, "/api/health") for name, port in [
        ("benue-street-locator", 3101), ("mindgrow", 3102), ("community-connect", 3103)
    ]]
    for name, folder, port, endpoint in previews:
        if not folder.is_dir():
            raise RuntimeError(f"Missing checkout: {folder}. Use --project-root to select it.")
        url = f"http://127.0.0.1:{port}"
        if not healthy(url, endpoint):
            if name == "portfolio":
                binary = STATE / "portfolio"
                if args.rebuild or not binary.exists():
                    subprocess.run(["go", "build", "-trimpath", "-o", str(binary), "."], cwd=folder, check=True)
                command = [str(binary)]
            else:
                if args.rebuild or not (folder / ".local-build/service.mjs").exists():
                    if not (folder / "node_modules").is_dir():
                        subprocess.run([npm, "ci"], cwd=folder, env=env, check=True)
                    subprocess.run([npm, "run", "build:local"], cwd=folder, env=env, check=True)
                command = [node, "scripts/local-server.mjs"]
            with (STATE / (name + ".log")).open("a") as log:
                # Normal previews always use the existing saved data, even if the
                # caller's shell previously configured an isolated QA directory.
                runtime_env = {key: value for key, value in env.items() if key != "APP_STATE_DIR"}
                process = subprocess.Popen(command, cwd=folder, env={**runtime_env, "PORT": str(port)},
                                           stdout=log, stderr=subprocess.STDOUT, start_new_session=True)
            processes[name] = {"pid": process.pid, "cwd": str(folder), "command": command}
            registry.write_text(json.dumps(processes, indent=2) + "\n")
            for _ in range(100):
                if process.poll() is not None:
                    raise RuntimeError(f"{name} exited. See {STATE / (name + '.log')}")
                if healthy(url, endpoint):
                    break
                time.sleep(.2)
            else:
                raise RuntimeError(f"{name} did not become ready. See {STATE / (name + '.log')}")
        print(name + ": " + url, flush=True)
        if not args.no_open:
            opener = shutil.which("xdg-open")
            if opener:
                subprocess.Popen([opener, url], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
            else:
                import webbrowser
                if not webbrowser.open_new_tab(url):
                    print("Open " + url + " in your browser.")


if __name__ == "__main__":
    try:
        main()
    except (RuntimeError, subprocess.CalledProcessError) as error:
        raise SystemExit(str(error))
