# Change

> yaml: change

Properties about changing the file

**Example:**

```yaml
...
    files:
        main:
            filepath: "src/main.c"
            change:
              ...
```

## Split On

> yaml: split-on
>
> String

Will split the file on specified string and will append after it.
If left empty appends at the end of the file.

_Uses strings.Split()_

**Example:**

```yaml
...
    files:
        main:
            filepath: "src/main.c"
            change:
                split-on = "MIDDLEWARE = [",
                ...
```

## Append

> yaml: append
>
> String

Content that will be appended after the split on

**Example:**

```yaml
    files:
        main:
            filepath: "src/main.c"
            change:
                split-on = "MIDDLEWARE = [",
                append = "\n\t'corsheaders.middleware.CorsMiddleware',"
```