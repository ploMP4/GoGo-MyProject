# GoGo MyProject

GoGo is a CLI tool that creates the starter boilerplate
for your projects and it's really helpful for people
who use many different programming languages and frameworks.

- [Installation](#installation)
  - [Linux](#linux)
    - [Release Branch](#1-using-the-release-branch)
    - [Makefile](#2-using-make)
    - [Install.sh](#3-using-the-installsh)
    - [Manual Install](#4-manual-installation)
    - [Adding command to PATH](#add-command-to-path)
    - [Installing Go and Make](#installing-go-and-make)
  - [Windows](#windows)
    - [Release Branch](#1-using-the-release-branch-1)
  - [Mac](#mac)
    - [Release Branch](#1-using-the-release-branch-2)
- [Usage](#usage)
- [Config Files](#config-files)
  - [Commands](#commands)
  - [Dirs](#dirs)
  - [Sub Commands](#sub-commands)
  - [Help](#help)
- [Templates](#templates)
- [Flags]()
  - [a, all]()
  - [e, exclude]()
  - [P, set-config-path]()
  - [h, help]()
  - [v, version]()
- [Dependencies](#dependencies)

---

## Installation

### Linux:

- #### 1. Using the release branch

- #### 2. Using Make

  Make sure you have make and go installed.  
  See [Installing Go and Make](#installing-go-and-make) if you need help installing them

  Clone the repo to your computer. Then cd into the project directory and run make install

  ```
  git clone https://github.com/ploMP4/GoGo-MyProject
  cd GoGo-MyProject && make install
  ```

  Finally [Add the command to PATH](#add-command-to-path)

- #### 3. Using the install.sh

  Make sure you have make and go installed.  
  See [Installing Go and Make](#installing-go-and-make) if you need help installing them

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

  Finally [Add the command to PATH](#add-command-to-path)

- #### 4. Manual installation

  Make sure you have make and go installed.  
  See [Installing Go and Make](#installing-go-and-make) if you need help installing them

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

  Create a new folder in your home directory and move the executable there in the `bin` folder. Then create a settings.json file

  ```
  mkdir $HOME/.gogo $HOME/.gogo/bin
  mv ./dist/gogo $HOME/.gogo/bin
  echo '{ "config-path": "" }' > $HOME/.gogo/bin/settings.json
  ```

  Finally [Add the command to PATH](#add-command-to-path)

- #### Add command to PATH

  - zsh

    Add this to your `.zshrc` file before the `export PATH` line

    ```zsh
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

  - bash

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

- #### Installing Go and Make

  - Debian:

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

  - Fedora:

    ```
    sudo dnf install make go
    ```

  - Arch:

    ```
    sudo pacman -S make go
    ```

### Windows:

- #### 1. Using the release branch

### Mac:

- #### 1. Using the release branch

---

## Usage

```
gogo <COMMAND> <APPNAME> [args]
```

---

## Config files

### Commands:

**_Example:_**

>Array\<Array\<string>>

Array with the commands that will be executed.
Commands should be passed as an array of strings 
instead of using spaces

```json
"commands": [
  ["npx", "create-react-app"]
]
```

### Dirs:

>Array\<string>

Directories with these names will be created at the root
of your project

**_Example:_**

```json
"dirs": ["src", "dist", "tests", "vendor"]
```

### Sub Commands:

- [Command Name](#command-name)
  - [Name](#name)
  - [Command](#command)
  - [Override](#override)
  - [Parallel](#parallel)
  - [Exclude](#exclude)
  - [Files](#files)
    - [Description](#description)
    - [Filepath](#filepath)
    - [Template](#template)
    - [Change](#change)
- [Help](#help-1)
- #### Command Name:
  >Object

  The key defined is the argument you need to pass
  to activate the subcommand. The value contains 
  data about what it does 

  **_Example:_**

  ```json
  "ts": {
    ...
  }
  ```

- #### Name:
  >String

  Name that will be displayed in status messages e.x Installing: Typescript

  **_Example:_**

  ```json
  "name": "Typescript"
  ```

- #### Command:
  >Array\<string>

  The command that will be executed.

  **_Example:_**

  ```json
  "command": ["npx", "create-react-app", "--template", "typescript"]
  ```

- #### Override:
  >Boolean
  >
  >Default: false

  If true the overrides the last command in the main
  commands array with this command.

  **_Example:_**

  ```json
  "override": true
  ```

- #### Parallel:
  >Boolean
  >
  >Default: false

  If true the command will be run concurrently with others

  **_Example:_**

  ```json
  "parallel": false
  ```

- #### Exclude:
  >Boolean
  >
  >Default: false

  If true this command will be ignored when the [a, all]() flag is used

  **_Example:_**

  ```json
  "exclude": false
  ```

- #### Files:

  - #### Description:

  - #### Filepath:

  - #### Template:

  - #### Change:

- #### Help:
  >String

  Help text for the command

  **_Example:_**

  ```json
  "help": "Use typescript template of cra"
  ```

### Help:

---

## Templates

---

## Dependencies

```
github.com/briandowns/spinner
github.com/fatih/color
```
