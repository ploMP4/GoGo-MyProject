---
sidebar_position: 2
---

# Create a template

A template is basically just a file that gets copied into your project.

So let's say we want to create a starter file for our c++ project to use
as a template. As stated before template files need to be under a directory
with the same name as the command.

Let's say we name our command cpp,
we will create a directory in the `templates folder` called `cpp`.

**When we say that a file is going to use a template, in our [config](../gadgets/sub-cmd/files),**
**we specify a path for that file so in our templates folder we need to follow the same directory structure.**

For example let's say we want to have a directory named `src` that will
store our source files, and in there we want to insert our `main.cpp template`.
Based on that we will need to create a directory with the name `src` inside
our template folder for the command we made called `cpp`.

```bash
mkdir src
```

Now we can create our template file inside the src directory.

```bash
cd src && touch main.cpp
```

```cpp title="templates/cpp/minimal/src/main.cpp"
#include <iostream>

int main() {
    std::cout << "Hello World" << std::endl;

    return 0;
}
```

Now we can add the template file under our desired subcommand like so:

```yaml title="cpp.yaml"
subCommands:
  cpp-subcommand:
    files:
      src:
        filepath: "src/main.cpp"
        template: true
```
