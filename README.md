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
  - [Quick start](#quick-start)
- [Built in flags](#flags)
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
  - [Get Started](#get-started)
  - [Creating a template](#creating-a-template)
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

  Create a new folder in your home directory and move the executable there in the `bin` folder. Then create a settings.toml file

  ```
  mkdir $HOME/.gogo $HOME/.gogo/bin
  mv ./dist/gogo $HOME/.gogo/bin
  echo '{ "config-path": "" }' > $HOME/.gogo/bin/settings.toml
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

### Quick Start:

---

## Flags

Built in flags of the app.

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

---

## Config files

Below is documentation for creating your own config
file. You can also use the [example](https://github.com/ploMP4/GoGo-MyProject/blob/main/examples/config/example.toml) file as a template
or modify the [already existing](https://github.com/ploMP4/GoGo-MyProject/tree/main/examples/config) ones.

### Commands:

> toml: "commands"
>
> Array\<Array\<string>>

**_Description:_**

Array with the commands that will be executed.
Commands should be passed as an array of strings
instead of using spaces.

**_Example:_**

```toml
commands = [
  ["npx", "create-react-app"]
]
```

### Dirs:

> toml: "dirs"
>
> Array\<string>

**_Description:_**

Directories with these names will be created at the root
of your project.

**_Example:_**

```toml
dirs = ["src", "dist", "tests", "vendor"]
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

  **_Description:_**

  This will be the name of the command when you run the app

  **_Example:_**

  ```toml
  [subCommands.ts]
  ...
  ```
  ```bash
  gogo react myapp ts
  ```

  - #### Name:

    > toml: "name"
    >
    > String

    **_Description:_**

    Name that will be displayed in log messages e.x Installing: Typescript.

    **_Example:_**

    ```toml
    [subCommands.ts]
    name = "Typescript"
    ```

  - #### Command:

    > toml: "command"
    >
    > Array\<string>

    **_Description:_**

    The command that will be executed.

    **_Example:_**

    ```toml
    command = ["npx", "create-react-app", "--template", "typescript"]
    ```

  - #### Override:

    > toml: "override"
    >
    > Boolean
    >
    > Default: false

    **_Description:_**

    If true the overrides the last command in the main
    commands array with this command.

    **_Example:_**

    ```toml
    override = true
    ```

  - #### Parallel:

    > toml: "parallel"
    >
    > Boolean
    >
    > Default: false

    **_Description:_**

    If true the command will be run concurrently with others.

    **_Example:_**

    ```toml
    parallel = false
    ```

  - #### Exclude:

    > toml: "exclude"
    >
    > Boolean
    >
    > Default: false

    **_Description:_**

    If true this command will be ignored when the [a, all](#a-all) flag is used.

    **_Example:_**

    ```toml
    exclude = false
    ```

  - #### Files:

    > toml: "files"

    **_Description:_**

    Specify files that you want to change

    - #### Description:

      **_Description:_**

      The string that will be displayed in log messages

      **_Example:_**

      ```toml
      [subCommands.cors.files.cors-middleware]
      ...
      ```

    - #### Filepath:

      > toml: "filepath"
      >
      > String

      **_Description:_**

      Path where the file we want to edit is located. **Path starts from the root file of our project**.

      **_Example:_**

      ```toml
      filepath = "src/main.c"
      ```

      You can also use the **\<APPNAME>** tag which searches for something
      with the same name as your app. Useful for things like a django project.

      **_Example:_**

      ```toml
      filepath = "<APPNAME>/settings.py"
      ```

    - #### Template:

      > toml: "template"
      >
      > Boolean
      >
      > Default: false

      **_Description:_**

      If true updates the file using a template. See [templates section](#templates) for more info.

      **_Example:_**

      ```toml
      template = false
      ```

    - #### Change:

      > toml: "change"

      **_Description:_**

      Properties about changing the file

      **_Example:_**

      ```toml
      change = {
          ...
       }
      ```

      - #### Split On

        > toml: "split-on"
        >
        > String

        **_Description:_**

        Will split the file on specified string and will append after it.
        If left empty appends at the end of the file.

        _Uses strings.Split()_

        **_Example:_**

        ```toml
        change = {
          split-on = "MIDDLEWARE = [",
        }
        ```

      - #### Append

        > toml: "append"
        >
        > String

        **_Description:_**

        Content that will be appended after the split on

        **_Example:_**

        ```toml
        change = {
          ...
          append = "\n\t'corsheaders.middleware.CorsMiddleware',"
        }
        ```

- #### Help:

  > toml: "help"
  >
  > String

  **_Description:_**

  Help text for the subcommand

  **_Example:_**

  ```toml
  help = "Use typescript template of cra"
  ```

### Help:

> toml: "help"
>
> String

**_Description:_**

Help text for the command

**_Example:_**

```toml
help = "Creates react app"
```

---

## Templates

Templates are pre-made files that you copy into your project.

### Get Started:

- First you need to create a directory where you will store your templates anywhere that you want in your machine.

```bash
mkdir ~/.gogo/templates
```

- Then you need you add it to `settings.toml` see [set-template-path](#t-set-template-path) for info on how to easily do that

- Finally you are ready to [create some templates](#creating-a-template).
  **Templates need to be stored in a folder with the same name as the command they are meant for.**

  For example if we have a command named `cpp` we need to store the
  templates for this command in `TEMPLATE_FOLDER/cpp`.

### Creating a template:

A template is basically just a file that gets copied into your project.

So let's say we want to create a starter file for our c++ project to use
as a template. As stated before template files need to be under a directory
with the same name as the command.

Let's say we name our command cpp,
we will create a directory in the `templates folder` called `cpp`.

**When we say that a file is going to use a template in our [config](#files) file**
**we specify a path for that file so in our templates folder we need to follow the same directory structure.**

For example let's say we want to have a directory named `src` that will
store our source files, and in there we want to insert our `main.cpp template`.
Based on that we will need to create a directory with the name `src` inside
our template folder for the command we made called `cpp`.

```bash
mkdir src
```

Now we can create our template file inside the src directory.

```bash
cd src && touch main.cpp
```

main.cpp

```cpp
#include <iostream>

int main() {
    std::cout << "Hello World" << std::endl;

    return 0;
}
```

Now we can add the template file under our desired subcommand like so:

```toml
[subCommands.cpp-subcommand.files.src]
filepath = "src/main.cpp"
template = true
```

---

## Dependencies

```
github.com/BurntSushi/toml
github.com/briandowns/spinner
github.com/fatih/color
github.com/mattn/go-colorable
github.com/mattn/go-isatty
golang.org/x/sys
```
