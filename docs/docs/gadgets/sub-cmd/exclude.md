---
sidebar_position: 6
---

# Exclude

> Default: false

If true this subcommand will be ignored when the [a, all](../../flags/all) flag is used.

**Example:**

```yaml
subCommands:
  ts:
    name: "Typescript"
    command: "npx create-react-app --template typescript"
    override: true
    exclude: true
```
