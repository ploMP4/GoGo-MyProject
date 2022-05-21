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
  - [Quick start]()
- [Flags](#flags)
  - [h, help](#h-help)
  - [v, version](#v-version)
  - [C, set-config-path](#c-set-config-path)
  - [T, set-template-path](#t-set-template-path)
  - [a, all](#a-all)
  - [e, exclude](#e-exclude)
- [Config Files](#config-files)
  - [Commands](#commands)
  - [Dirs](#dirs)
  - [Sub Commands](#sub-commands)
  - [Help](#help)
- [Templates](#templates)
  - [Get Started]()
  - [Pros & Cons]()
  - [Creating a template]()
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

## Flags

### h, help:

> gogo help [command]

**_Description:_**

Displays the help menu. If a command name is passed as an argument
displays the help menu for the specific command

**_Example:_**

```bash
gogo help
```

```bash
gogo help react
```

OR

```bash
gogo h
```

```bash
gogo h react
```

### v, version:

**_Description:_**

Show the application version

**_Example:_**

```bash
gogo version
```

OR

```bash
gogo v
```

### C, set-config-path:

**_Description:_**

Set the path of the folder containing the config files.

**NOTE: You need to specify the full path from the root directory**

**_Example:_**

```bash
gogo set-config-path /home/me/.gogo/config
```

OR

```powershell
gogo C C:\Users\Me\gogo\config
```

### T, set-template-path:

**_Description:_**

Set the path of the folder containing the templates.

**NOTE: You need to specify the full path from the root directory**

**_Example:_**

```bash
gogo set-template-path /home/me/.gogo/templates
```

OR

```powershell
gogo T C:\Users\Me\gogo\templates
```

### a, all:

**_Description:_**

Runs all the subcommands for the command except the ones
that have the exclude value set to true.

**_Example:_**

```bash
gogo all react
```

OR

```bash
gogo a react
```

### e, exclude:

> gogo exclude [subcommand]

**_Description:_**

**_Example:_**

---

## Config files

Below is documentation for creating your own config
file. You can also use the [example](https://github.com/ploMP4/GoGo-MyProject/blob/main/examples/config/example.json) file as a template
or modify the [already existing](https://github.com/ploMP4/GoGo-MyProject/tree/main/examples/config) ones.

### Commands:

> json: "commands"
>
> Array\<Array\<string>>

**_Description:_**

Array with the commands that will be executed.
Commands should be passed as an array of strings
instead of using spaces.

**_Example:_**

```json
"commands": [
  ["npx", "create-react-app"]
]
```

### Dirs:

> json: "dirs"
>
> Array\<string>

**_Description:_**

Directories with these names will be created at the root
of your project.

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

  > Object

  **_Description:_**

  The key defined is the argument you need to pass
  to activate the subcommand. The value contains
  data about what it does.

  **_Example:_**

  ```json
  "ts": {
    ...
  }
  ```

  - #### Name:

    > json: "name"
    >
    > String

    **_Description:_**

    Name that will be displayed in status messages e.x Installing: Typescript.

    **_Example:_**

    ```json
    "name": "Typescript"
    ```

  - #### Command:

    > json: "command"
    >
    > Array\<string>

    **_Description:_**

    The command that will be executed.

    **_Example:_**

    ```json
    "command": ["npx", "create-react-app", "--template", "typescript"]
    ```

  - #### Override:

    > json: "override"
    >
    > Boolean
    >
    > Default: false

    **_Description:_**

    If true the overrides the last command in the main
    commands array with this command.

    **_Example:_**

    ```json
    "override": true
    ```

  - #### Parallel:

    > json: "parallel"
    >
    > Boolean
    >
    > Default: false

    **_Description:_**

    If true the command will be run concurrently with others.

    **_Example:_**

    ```json
    "parallel": false
    ```

  - #### Exclude:

    > json: "exclude"
    >
    > Boolean
    >
    > Default: false

    **_Description:_**

    If true this command will be ignored when the [a, all]() flag is used.

    **_Example:_**

    ```json
    "exclude": false
    ```

  - #### Files:

    > json: "files"
    >
    > Object

    **_Description:_**

    Specify files that you want to change

    - #### Description:

      > Object

      **_Description:_**

      The key of the object is just what is going
      to be shown in the message when executing.
      The value contains data about what it does.

      **_Example:_**

      ```json
      "CORS middleware": {
          ...
       },
      ```

    - #### Filepath:

      > json: "filepath"
      >
      > String

      **_Description:_**

      Path where the file we want to edit is located. **Path starts from the root file of our project**.

      **_Example:_**

      ```json
      "filepath": "src/main.c"
      ```

      You can also use the **\<APPNAME>** tag which searches for a something
      with the same name as your app. Useful for things like a django project.

      **_Example:_**

      ```json
      "filepath": "<APPNAME>/settings.py"
      ```

    - #### Template:

      > json: "template"
      >
      > Boolean
      >
      > Default: false

      **_Description:_**

      If true updates the file using a template. See [Creating a template]() for more info.

      **_Example:_**

      ```json
      "template": false
      ```

    - #### Change:

      > json: "change"
      >
      > Object

      **_Description:_**

      Properties about changing the file

      **_Example:_**

      ```json
      "change": {
          ...
       }
      ```

      - #### Split On

        > json: "split-on"
        >
        > String

        **_Description:_**

        Will split the file on specified string and will append after it.
        If left empty appends at the end of the file.

        _Uses strings.Split()_

        **_Example:_**

        ```json
        "split-on": "MIDDLEWARE = [",
        ```

      - #### Append

        > json: "append"
        >
        > String

        **_Description:_**

        Content that will be appended after the split on

        **_Example:_**

        ```json
        "append": "\n\t'corsheaders.middleware.CorsMiddleware',"
        ```

- #### Help:

  > json: "help"
  >
  > String

  **_Description:_**

  Help text for the subcommand

  **_Example:_**

  ```json
  "help": "Use typescript template of cra"
  ```

### Help:

> json: "help"
>
> String

**_Description:_**

Help text for the command

**_Example:_**

```json
"help": "Creates react app"
```

---

## Templates

---

## Dependencies

```
github.com/briandowns/spinner
github.com/fatih/color
github.com/mattn/go-colorable
github.com/mattn/go-isatty
golang.org/x/sys
```
