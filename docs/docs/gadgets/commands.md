---
sidebar_position: 1
---

# Commands

Array with the commands that will be executed.
The last command in the array gets passed the appname
as the last argument.

**Example:**

```yaml
commands:
  - echo "this is a command"
  - echo "this is another command"
  - npx create-react-app
```
