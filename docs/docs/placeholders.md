---
sidebar_position: 6
---

# Built in Placeholders

Coming soon...

## \_APPNAME

The `_APPNAME` placeholder is used inside the `gadget` files in the `commands`
and `files` properties and will be replaced with whatever we pass as an argument.

**Example:**

We define a `gadget` that creates a rust project using `cargo new` to have
a dynamic project name we pass the \_APPNAME placeholder at the end.

```yaml title="~/.gogo/gadgets/rust.yaml"
commands: ["cargo new _APPNAME"]
```

Now when we use the gadget we give a value for the `_APPNAME` and
wherever there is \_APPNAME inside the commands array it will be replaced
with the value we gave it.

```bash
gogo rust _APPNAME rewriteitinrust
```

## \_FILENAME

Coming soon...
