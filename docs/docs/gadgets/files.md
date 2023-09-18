---
sidebar_position: 2
---

# Files

The files property gives us the ability to modify files or create them from prebuilt
templates.

## Description

The yaml key will be used as a description to be displayed in log messages

**Example:**

```yaml title="~/.gogo/gadgets/expressjs.yaml"
files:
  cors-middleware: ...
```

**Example Output:**

```
Adding: [cors-middleware in src/app.js]
```

## Filepath

Path where the file we want to edit is located. **Path starts from the root of our project**.

**Example:**

```yaml
files:
  cors-middleware:
    filepath: "src/app.js"
```

You can also use the `_APPNAME` placeholder which searches for a directory
with the same name as your appname. Useful for things like a django project.

**Example:**

```yaml
files:
  settings_py:
    filepath: "_APPNAME/settings.py"
```

## Template

> Default: false

If true updates the file using a template. See [templates section](../templates) for more info.

**Example:**

```yaml
files:
  app.js:
    filepath: "src/app.js"
    template: true
```

## Change

Properties about changing the file

**Example:**

```yaml
files:
  app.js:
    filepath: "src/app.js"
    change: ...
```

### Split On

Will split the file on specified string and will append after it.
If left empty appends at the end of the file.

_Uses golang strings.Split() under the hood_

**Example:**

```yaml
files:
  cors-middleware:
    filepath: "_APPNAME/settings.py"
    change:
        split-on: "MIDDLEWARE = [",
        ...
```

### Append

Content that will be appended after the [split on](#split-on)

**Example:**

```yaml
files:
  cors-middleware:
    filepath: "_APPNAME/settings.py"
    change:
        split-on: "MIDDLEWARE = [",
        append: "\n\t'corsheaders.middleware.CorsMiddleware',"
```

### Placeholder

We can also define placeholder strings that exist inside our file that we want to edit and replace them either with
a default value that we provide in the config or by a custom value that we pass as an argument.

**Example:**

```yaml title="gadgets/react-component.yaml"
files:
  component:
    filepath: "components/component.jsx"
    template: true
    change:
      placeholder:
        _NAME: "Component"
```

Let's say we are creating a react function component from a template, now to any dynamic parts that have
to do with the components name we can give the value `_NAME`.

```jsx title="templates/components/component.jsx"
import React from "react";

const _NAME = () => {
  return <div>_NAME</div>;
};

export default _NAME;
```

Now we can run our gadget and provide a value for the `_NAME` placeholder like so:

```bash
gogo react-component _NAME Button
```

Our newly created file will look like this:

```jsx title="components/component.jsx"
import React from "react";

const Button = () => {
  return <div>Button</div>;
};

export default Button;
```

_If no value is provided when we run the command the placeholder will take the default value defined in our gadget._
_For this example it will have the value `Component`_

_Usually placeholders start with an underscore (\_) followed by the name of the placeholder in uppercase._
_But really they can be any string we want_

**Example of a different placeholder name:**

```yaml
change:
  placeholder:
    _NAME: "Component"
```

There are also some built in placeholders that you can find [here](../../placeholders).
