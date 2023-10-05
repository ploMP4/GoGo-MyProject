---
sidebar_position: 1
---

# GoGo My Intro

GoGo is an extendable cli tool that allows you to easily blueprint boilerplate and repetitive commands.
From basic files to scaffolding the entire project structure.

## How it works

GoGo utilizes something called `gadgets`, these are just `yaml` config files that describe what each `gadget` does.
`Gadgets` can be invoked using the `gogo` command and are like small cli applications themselves. They can run shell commands,
modify and copy files from prebuilt templates and they can also have their own `subcommands` that are capable of the same things aswell.

## Getting Started

Get started by **[Installing gogo](./category/installation/)**.

Make sure that gogo is installed by running the following:

```bash
gogo version
```

## Examples

After you [installed](./category/installation) gogo, you can copy
some example gadgets and templates from the [examples](https://github.com/ploMP4/GoGo-MyProject/tree/main/examples) folder in the repo. After that place the `gadgets` and `templates` folders inside your `.gogo` directory. Now you should see the example gadgets on the help menu and be able to use them.

```bash
gogo help
```
