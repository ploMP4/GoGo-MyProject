---
sidebar_position: 2
---

# Name

The name property will be displayed in a log message when we use a subcommand.

**Example:**

Let's say we have the subcommand `ts` that adds typescript support when we create
a react app. We give it the name `Typescript` so when we see the log message we know
that we included typescript in our project.

```yaml title="gadgets/react.yaml"
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
