---
sidebar_position: 2
---

# Filepath

Path where the file we want to edit is located. **Path starts from the root of our project**.

**Example:**

```yaml
files:
  main:
    filepath: "src/app.js"
```

You can also use the **\_APPNAME** tag which searches for something
with the same name as your app. Useful for things like a django project.

**Example:**

```yaml
files:
  settings_py:
    filepath: "_APPNAME/settings.py"
```
