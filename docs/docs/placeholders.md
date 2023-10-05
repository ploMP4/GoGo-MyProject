---
sidebar_position: 6
---

# Built in Placeholders

GoGo has some built in placeholders that you can use to help you make your `gadgets`
more dynamic.

## \_APPNAME

The `_APPNAME` placeholder can be used inside the `gadget` files in the `commands`
and `files` properties and will be replaced with whatever we pass as an argument to it.

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

The `_FILENAME` placeholder can be used when we call a `gadget` that has a single
file defined in the `files` property and will rename that file to whatever we pass as
an argument to it.

**_Note: if there is more than one file in defined in the `files` property of the gadget
the `_FILENAME` placeholder will be ignored by `gogo`._**

**Example:**

`_FILENAME` can be usefull when we have a `gadget` that generates a single file with some boilerplate
inside like a react component.

The `gadget` will look like this.

```yaml title=".gogo/gadgets/component.yaml"
template: true
help: "Create a react function component"

files:
  component:
    filepath: "components/component.jsx"
    template: true
    change:
      placeholder:
        _NAME: "Component"
```

When we run `gogo component` it will generate a new react component inside the components directory
containing the code that we've included in the template. But we don't really want our newly created file
to be named component.jsx, we can overcome that by using the `_FILENAME` placeholder to rename the file to
whatever we want.

```
gogo component _FILENAME button.jsx
```

By running the command above we will now generate a file called button.jsx inside the components directory.

**_NOTE: We need to specify the file extention aswell, providing only a name will remove the files extention._**
