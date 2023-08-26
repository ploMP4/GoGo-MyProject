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

DIST_PATH = Path(f"{Path.cwd()}/dist")
BINARY_PATH = Path(f"{APPLICATION_PATH}/bin")
SETTINGS_FILE_PATH = Path(f"{BINARY_PATH}/settings.yaml")
GADGET_PATH = Path(f"{APPLICATION_PATH}/gadgets")
TEMPLATES_PATH = Path(f"{APPLICATION_PATH}/templates")

if __name__ == "__main__":
    subprocess.run("go build -o ./dist/gogo ./cmd/...", shell=True)

    if not os.path.isdir(APPLICATION_PATH):
        os.mkdir(APPLICATION_PATH)
        os.mkdir(BINARY_PATH)
        os.mkdir(GADGET_PATH)
        os.mkdir(TEMPLATES_PATH)

    shutil.move(DIST_PATH, BINARY_PATH)

    with open(SETTINGS_FILE_PATH, "a") as f:
        f.write(f'gadget-path: "{APPLICATION_PATH}/gadgets"\n')
        f.write(f'template-path: "{APPLICATION_PATH}/templates"')
