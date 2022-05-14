# GoGo MyProject

GoGo is a CLI tool that creates the starter boilerplate 
for your projects and it's really helpful for people
who use many different programming languages and frameworks.

- [Installation](#installation)
  - [Linux](#linux)
    - [Makefile](#1-using-make)
    - [Install.sh](#2-using-the-installsh)
    - [Manual Install](#3-manual-installation)
    - [Adding command to PATH](#add-command-to-path)
    - [Installing Go and Make](#installing-go-and-make)
  - [Windows](#windows)
  - [Mac](#mac)
- [Usage](#usage)
- [Config Files](#config-files)
- [Templates]()

---

## Installation:

### Linux:

- #### 1. Using Make:

  Make sure you have make and go installed.  
  See [Installing Go and Make](#installing-go-and-make) if you need help installing them

  Clone the repo to your computer. Then cd into the project directory and run make install

  ```
  git clone https://github.com/ploMP4/GoGo-MyProject
  cd GoGo-MyProject && make install
  ```

  Finally [Add the command to PATH](#add-command-to-path)

- #### 2. Using the install.sh

  Make sure you have make and go installed.  
  See [Installing Go and Make](#installing-go-and-make) if you need help installing them

  Clone the repo to your computer:

  `git clone https://github.com/ploMP4/GoGo-MyProject`

  cd into the project directory and build the project using go build.

  `cd GoGo-MyProject && go build -o ./dist/gogo ./cmd/...`

  OR
  
  `cd GoGo-MyProject && make build`

  Run the install.sh script

  `./scripts/install.sh`

  Finally [Add the command to PATH](#add-command-to-path)

- #### 3. Manual installation

  Make sure you have make and go installed.  
  See [Installing Go and Make](#installing-go-and-make) if you need help installing them

  Clone the repo to your computer:

  >`git clone https://github.com/ploMP4/GoGo-MyProject`

  cd into the project directory and build the project using go build.

  `cd GoGo-MyProject && go build -o ./dist/gogo ./cmd/...`

  OR

  `cd GoGo-MyProject && make build`

  Create a new folder in your home directory and move the executable there in the *bin* folder. Then create a settings.json file


  ```
  mkdir $HOME/.gogo $HOME/.gogo/bin
  mv ./dist/gogo $HOME/.gogo/bin
  echo '{ "config-path": "" }' > $HOME/.gogo/bin/settings.json
  ```

  Finally [Add the command to PATH](#add-command-to-path)


- #### Add command to PATH
  - zsh

    Add this to your .zshrc file before the `export PATH` line 

    `path+=($HOME'/.gogo/bin')`

    *if you installed it in a different directory add that instead*

  - bash

    Add this to the end of your .bashrc file

    `export PATH="$HOME/.gogo/bin:$PATH"`

- #### Installing Go and Make

  - Debian:

    `sudo apt install make`

    Visit the [go website](https://go.dev/dl/) and download the linux version.
    Then unzip the file you downloaded.

    `sudo tar -C /usr/local -xzf <filename>`

    Finally add this line to your .bashrc file

    `export PATH="/usr/local/go/bin:$PATH"`

  - Fedora:  

    `sudo dnf install make go`

  - Arch:

    `sudo pacman -S make go`

### Windows:

### Mac:

---

## Usage:

`gogo <COMMAND> <APPNAME> [args]`

---

## Config files:

