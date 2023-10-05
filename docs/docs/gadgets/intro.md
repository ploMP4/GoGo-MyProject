---
sidebar_position: 1
---

# Get Started

To create a `gadget` we need to start by creating a `yaml` file inside the gadgets
directory.

```bash title="/home/exampleuser/.gogo/gadgets/"
touch react.yaml
```

Now we are going to add functionallity to our newly created `gadget` by setting some properties.

## Commands

The commands property is an array with the commands that will be executed.

If we are using a command like `create-react-app` under the hood, we can use
the `_APPNAME` placeholder to be able to have a dynamic appname when we create our project.

We will see more about placeholders later in this guide.

**Example:**

```yaml title="~/.gogo/gadgets/react.yaml"
commands:
  - echo "this is a command"
  - echo "this is another command"
  - npx create-react-app _APPNAME
```

## Chdir

> Default: false

The `chdir` property is a boolean property that when `true` it will
try to change directory and go inside the one that matches the `_APPNAME` placeholder.
The rest of the program is going to be executed from there.

```yaml title="~/.gogo/gadgets/react.yaml"
commands: ["npx create-react-app _APPNAME"]
chdir: true
```

## Dirs

This property just creates the directories with the names
specified in the array.

_If we have the [chdir](#chdir) property set to true the directories will be created inside the folder that matches the `_APPNAME` placeholder_

**Example:**

```yaml title="~/.gogo/gadgets/react.yaml"
commands: ["npx create-react-app _APPNAME"]
chdir: true
dirs: ["components", "layouts"]
```

## Help

Help text that will be displayed in the `gogo` help menu.

**Example:**

```yaml
help: "Create a react project"
```
