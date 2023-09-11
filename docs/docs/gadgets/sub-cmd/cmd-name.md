---
sidebar_position: 1
---

# Command Name

The command name is the key that you specify inside the yaml gadget file. this key will be the name of the argument
for when you want to use it as a subcommand.

**Example:**

Let's say we have a gadget called `react` that creates a react app and sometimes we want to add
typescript support to our app. Because of that we create a subcommand to add typescript support whenever we want,
so under the `subCommands` section we add the subcommand with the key `ts` for typescript.

```yaml title="gadgets/react.yaml"
subCommands:
  ts: ...
```

After we write our configuration we can now use the `ts` key as an argument to the react gadget
to add typescript support to our project.

```bash
gogo react myappname ts
```
