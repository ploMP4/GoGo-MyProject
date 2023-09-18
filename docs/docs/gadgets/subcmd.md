---
sidebar_position: 3
---

# Sub Commands

## Command Name

The command name is the key that you specify inside the yaml gadget file. The same name as the key will also have the argument
for when you want to use it as a subcommand.

**Example:**

Let's say we have a gadget called `react` that creates a react app and sometimes we want to add
typescript support to our app. Because of that we create a subcommand to add typescript support whenever we want,
so under the `subCommands` section we add the subcommand with the key `ts` for typescript.

```yaml title="~/.gogo/gadgets/react.yaml"
subCommands:
  ts: ...
```

After we write our configuration we can now use the `ts` key as an argument to the react gadget
to add typescript support to our project.

```bash
gogo react myappname ts
```

## Name

The name property will be displayed in a log message whenever we use a subcommand.

**Example:**

Following the previous we have the subcommand `ts` that adds typescript support when we create
a react app. We give it the name `Typescript` so when we see the log message we know
that we included typescript in our project.

```yaml title="~/.gogo/gadgets/react.yaml"
subCommands:
  ts:
    name: "Typescript"
```

When we run the command with the `ts` argument the output should look like this:

Command:

```
gogo react myreactapp ts
```

Output:

```
Using: [Typescript]
Running: [npx create-react-app --template typescript myreactapp]
Created: [myreactapp]
```

## Commands

The commands property is an array with the commands that will be executed.

**Example:**

```yaml title="~/.gogo/gadgets/react.yaml"
subCommands:
  ts:
    name: "Typescript"
    commands: ["npx create-react-app --template typescript"]
```

## Override

> Default: false

When the override property is true in a subcommand the `commands` of that specific
subcommand will override the main gadget `commands` array.

**Example:**

Let's say we have a gadget that creates a react app using the `create-react-app` cli
and we occasionally want to use the typescript template. To satisfy these requirements
we will create a `gadget` that uses the normal `create-react-app` by default and a subcommand
that will `override` the default `create-react-app` and replace it with the typescript version instead.

```yaml title="~/.gogo/gadgets/react.yaml"
commands: ["npx create-react-app"]

subCommands:
  ts:
    name: "Typescript"
    commands: ["npx create-react-app --template typescript"]
    override: true
```

## Parallel

> Default: false

The subcommands that have the `parallel` property set to `true` will be run concurrently.

**Warning:**

Make sure that your `parallel` subcommands will not modify the same things or else
you might run in some race conditions.

**Example:**

```yaml title="~/.gogo/gadgets/react.yaml"
subCommands:
    tailwind:
        ...
        parallel: true

    jest:
        ...
        parallel: true
```

## Exclude

> Default: false

If true this subcommand will be ignored when the [a, all](../flags.mdx#a-all) flag is used.

**Example:**

```yaml title="~/.gogo/gadgets/react.yaml"
subCommands:
  ts:
    name: "Typescript"
    command: "npx create-react-app --template typescript"
    override: true
    exclude: true
```

## Help

Help text for the subcommand

**Example:**

```yaml
help: "Use typescript version of create-react-app"
```

## Files

Files work in the exact same way as the main gadget files, documentation can me found here [files](./files.md)

**Example:**

```yaml
subCommands:
  ts:
    name: ...
    command: ...
    files:
      App.tsx: ...
```
