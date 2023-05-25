---
sidebar_position: 1
---

# Linux

Installation instructions for linux opperating systems

## Download binary

Comming soon...

## Build from source (Makefile)

Make sure you have [Make and go installed](#installing-go-and-make).

Clone the repo to your computer. Then cd into the project directory and run `make install`

```
git clone https://github.com/ploMP4/GoGo-MyProject
cd GoGo-MyProject && make install
```

Finally [add the command to PATH](#add-command-to-path) to be able to execute it globally

## Using install.sh 

Make sure you have [Make and go installed](#installing-go-and-make).

Clone the repo to your computer:

```
git clone https://github.com/ploMP4/GoGo-MyProject
```

cd into the project directory and build the application.

```
cd GoGo-MyProject && go build -o ./dist/gogo ./cmd/...
```

OR

```
cd GoGo-MyProject && make build
```

Run the install.sh script

```
./scripts/install.sh
```

Finally [add the command to PATH](#add-command-to-path) to be able to execute it globally

## Adding command to PATH

### zsh

Add this to your `.zshrc` file before the `export PATH` line

```bash
path+=($HOME'/.gogo/bin')
```

_if you installed it in a different directory add that instead_

Restart your terminal or run to following command

```
source .zshrc
```

Test that everything went well by running

```
gogo version
```

### bash

Add this to the end of your `.bashrc` file

```bash
export PATH="$HOME/.gogo/bin:$PATH"
```

_if you installed it in a different directory add that instead_

Restart your terminal or run to following command

```
source .bashrc
```

Test that everything went well by running

```
gogo version
```

## Installing Go and Make

### Debian

```
sudo apt install make
```

Visit the [go website](https://go.dev/dl/) and download the linux version.
Then unzip the file you downloaded.

```
sudo tar -C /usr/local -xzf <filename>
```

Finally add this line to your .bashrc file

```bash
export PATH="/usr/local/go/bin:$PATH"
```

### Fedora

```
sudo dnf install make go
```

### Arch

```
sudo pacman -S make go
```
