---
sidebar_position: 7
---

# Local .gogo directory

Many times when we are working on a big project we wish we had a small, easily extendable CLI tool
to help us run repetitive commands and create/edit files. We can create something similar by making a `.gogo`
directory at the root of our project.

When we execute `gogo` it doesn't only look for `gadgets` inside our `gadgets/` directory but it also looks for
a `.gogo` directory existing where the command was run from.

Following the [\_FILENAME](./placeholders.md#_filename) example
where we created a `gadget` just for making a react component, we would want to place that inside a `.gogo` directory at the root
of our react project instead of having it as a globally accessible command.

```yaml title="myreactapp/.gogo/gadgets/component.yaml
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

```jsx title="myreactapp/.gogo/templates/component/components/component.jsx
import React from "react";

const _NAME = () => {
  return <div>_NAME</div>;
};

export default _NAME;
```

_`.gogo` directories can have a `gadgets/` and a `templates/` directory where we would define our gadgets and templates respectively._

Now we can navigate inside `myreactapp` and run `gogo component` from there to create a react component inside our components directory defined
within our application.

```
gogo component _FILENAME button.tsx _NAME Button
```
