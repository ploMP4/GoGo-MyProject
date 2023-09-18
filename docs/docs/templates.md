---
sidebar_position: 5
---

# Templates

Templates are pre-made files that you copy into your project.

## Getting started

First you need to create a directory where you will store your templates anywhere that you want in your machine.

_We usually store our templates under the .gogo/templates directory._

```bash
mkdir ~/.gogo/templates
```

Then you need you add it to `settings.yaml` see [templatedir](./flags.mdx#t-templatedir) for info on how to easily do that

**Templates need to be stored in a folder with the same name as the command they are meant for.**

For example if we have a command named `cpp` we need to store the templates for this command in `~/.gogo/templates/cpp`.

## Create a template

A template is basically just a file that gets copied into your project.

So let's say we want to create a starter file for our c++ project to use
as a template. As stated before template files need to be under a directory
with the same name as the command.

Let's say we name our command cpp,
we will create a directory in the `templates folder` called `cpp`.

**When we say that a file is going to use a template, in our [gadget](./gadgets/files.md#template),**
**we specify a path for that file so in our templates folder we need to follow the same directory structure.**

For example let's say we want to have a directory named `src` that will
store our source files, and in there we want to insert our `main.cpp template`.
Based on that we will need to create a directory with the name `src` inside
our `template folder` for the command we made called `cpp`.

```bash title="~/.gogo/templates/cpp/"
mkdir src
```

_If the template is used from a subcommand and not from the global gadget scope
we also need to create a directory with the same name as the subcommand. For example if
we have a subcommand called opengl that calls this template the same template would be located in
`~/.gogo/templates/cpp/opengl/src/main.cpp`_

Now we can create our template file inside the src directory.

```bash
cd src && touch main.cpp
```

```cpp title="~/.gogo/templates/cpp/src/main.cpp"
#include <iostream>

int main() {
    std::cout << "Hello World" << std::endl;

    return 0;
}
```

Now we can add the template file like so:

```yaml title="~/.gogo/gadgets/cpp.yaml"
files:
  src:
    filepath: "src/main.cpp"
    template: true
```
