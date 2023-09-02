#!/bin/python

import os
import shutil
import sys
import subprocess
from pathlib import Path

if sys.platform.startswith("win"):
    APPLICATION_PATH = Path(f"{Path.home()}/gogo")
else:
    APPLICATION_PATH = Path(f"{Path.home()}/.gogo")

DIST_PATH = Path(f"{Path.cwd()}/dist/gogo")
BINARY_PATH = Path(f"{APPLICATION_PATH}/bin")
SETTINGS_FILE_PATH = Path(f"{BINARY_PATH}/settings.yaml")
GADGET_PATH = Path(f"{APPLICATION_PATH}/gadgets")
TEMPLATES_PATH = Path(f"{APPLICATION_PATH}/templates")

if __name__ == "__main__":
    if sys.platform.startswith("win"):
        subprocess.run("go build -o ./dist/gogo.exe ./cmd/...", shell=True)
    else:
        subprocess.run("go build -o ./dist/gogo ./cmd/...", shell=True)

    if not os.path.isdir(APPLICATION_PATH):
        os.mkdir(APPLICATION_PATH)
        os.mkdir(BINARY_PATH)
        os.mkdir(GADGET_PATH)
        os.mkdir(TEMPLATES_PATH)

    if sys.platform.startswith("win"):
        shutil.move(DIST_PATH, BINARY_PATH / "gogo.exe")
    else:
        shutil.move(DIST_PATH, BINARY_PATH / "gogo")

    if not os.path.isfile(SETTINGS_FILE_PATH):
        with open(SETTINGS_FILE_PATH, "a") as f:
            f.write(f'gadget-path: "{APPLICATION_PATH}/gadgets"\n')
            f.write(f'template-path: "{APPLICATION_PATH}/templates"')
