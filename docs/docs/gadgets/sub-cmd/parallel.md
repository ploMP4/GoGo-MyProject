---
sidebar_position: 5
---

# Parallel

> Default: false

The subcommands that have the `parallel` property set to `true` will be run concurrently.

**Warning:**

Make sure that your `parallel` subcommands will not modify the same things or else
you might have some race conditions.

**Example:**

```yaml
subCommands:
    cors:
        ...
        parallel: true

    jwt:
        ...
        parallel: true
```
