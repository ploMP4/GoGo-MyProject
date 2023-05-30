---
sidebar_position: 2
---

# Filepath

> yaml: filepath
>
> String

Path where the file we want to edit is located. **Path starts from the root file of our project**.

**Example:**

```yaml
...
    files:
        main:
            filepath: "src/main.c"
```

You can also use the **<APPNAME\>** tag which searches for something
with the same name as your app. Useful for things like a django project.

**Example:**

```yaml
...
    files:
        settings_py:
            filepath: "<APPNAME>/settings.py"
```
