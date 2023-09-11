---
sidebar_position: 4
---

# Override

> Default: false

When the override property is true in a subcommand the `command` of that specific
subcommand will override the last value in the main gadget `commands` array.

**Example:**

Let's say we have a gadget that creates a react app using the `create-react-app` cli
and we occasionally want to use the typescript template. To satisfy these requirements
we will create a `gadget` that uses the normal `create-react-app` by default and a subcommand
that will `override` the default `create-react-app` and replace it with the typescript version instead.

```yaml title="gadgets/react.yaml"
commands: ["npx create-react-app"]

subCommands:
  ts:
    name: "Typescript"
    command: "npx create-react-app --template typescript"
    override: true
```
