---
sidebar_position: 3
---

# Sub Commands

## Command Name

This will be the name of the command when you run the app

**Example:**

```toml
[subCommands.ts]
...
```
```bash
gogo react myapp ts
```

### Name

> toml: "name"
>
> String

Name that will be displayed in log messages e.x Installing: Typescript.

**Example:**

```toml
[subCommands.ts]
name = "Typescript"
```

### Command

> toml: "command"
>
> Array<string\>

**_Description:_**

The command that will be executed.

**_Example:_**

```toml
command = ["npx", "create-react-app", "--template", "typescript"]
```

### Override

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

### Parallel

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

### Exclude

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

### Files

> toml: "files"

**_Description:_**

Specify files that you want to change

##### Description

The string that will be displayed in log messages

**_Example:_**

```toml
[subCommands.cors.files.cors-middleware]
...
```

##### Filepath

> toml: "filepath"
>
> String

**_Description:_**

Path where the file we want to edit is located. **Path starts from the root file of our project**.

**_Example:_**

```toml
filepath = "src/main.c"
```

You can also use the **<APPNAME\>** tag which searches for something
with the same name as your app. Useful for things like a django project.

**_Example:_**

```toml
filepath = "<APPNAME>/settings.py"
```

##### Template:

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

##### Change:

> toml: "change"

**_Description:_**

Properties about changing the file

**_Example:_**

```toml
change = {
  ...
}
```

##### Split On

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

##### Append

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

## Help

> toml: "help"
>
> String

**_Description:_**

Help text for the subcommand

**_Example:_**

```toml
help = "Use typescript template of cra"
```
