---
sidebar_position: 2
---

# Windows

Installation instructinons for windows opperating systems

## What you'll need

- Go >= 1.20
- Python >= 3.11

## Using the install script

Clone the repo to your computer.

cd into the project directory and run the `install` script

```
git clone https://github.com/ploMP4/GoGo-MyProject
cd GoGo-MyProject
python scripts\install.py
```

Finally [add the command to PATH](#add-command-to-path) to be able to execute it globally.

## Add command to PATH

### Using the command line

Run the Windows Command Prompt by typing `cmd` in the `Startup` menu.

To permanently set a directory into the Path environment variable settings, run the `setx` command:

```cmd
setx path "%PATH%;C:\Users\[username]\gogo\bin"
```

_Replace [username] with your users name._

### Using GUI

Open the Environment Variables settings by typing `Environment Variables` in the `Startup` menu and selecting the `Edit environment variables for your account` option.

In the opened window press the `Environment Variables` button that is located at the bottom.

Choose the `Path` option from the `System variables` panel and press the `Edit` button.

Then, click the `New` button to add the path to your gogo installation. _Should be (C:\Users\\[username]\gogo\bin)_

_Replace [username] with your users name._
