---
sidebar_position: 1
---

# Linux

Installation instructions for linux opperating systems

## Build from source (Recommended)

### What you'll need

- Go >= 1.20

### Using Make

Clone the repo to your computer. Then cd into the project directory and run `make install`

```
git clone https://github.com/ploMP4/GoGo-MyProject
cd GoGo-MyProject
make install
```

Finally [add the command to PATH](#add-command-to-path) to be able to execute it globally.

### Using the install.sh script

Clone the repo to your computer:

```
git clone https://github.com/ploMP4/GoGo-MyProject
```

cd into the project directory and run the `install` script

```
cd GoGo-MyProject
./scripts/install.sh
```

Finally [add the command to PATH](#add-command-to-path) to be able to execute it globally.

## Precompiled binary

Download the precompiled binary for your operating system from the [releases page](https://github.com/ploMP4/GoGo-MyProject/releases).

## Add command to PATH

Add this to the end of your `.bashrc` or `.zshrc` file

```bash
export PATH="$HOME/.gogo/bin:$PATH"
```

_if you installed it in a different directory, change `.gogo` to the directory you installed it instead._

Restart your terminal or run the following command

For bash users

```bash
source ~/.bashrc
```

For zsh users

```bash
source ~/.zshrc
```

Verify that everything went well by running

```bash
gogo version
```
