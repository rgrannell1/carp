
# carp 🐟

carp is a simple dependency-checker that checks a host matches the expected configuration. It performs dependency-checks in parallel, so it's very fast.

## Usage

```zsh
carp --file <fpath>
```

## Stability Index

1, Experimental - This project might die, it's undertested and underdocumented, and redesigns and breaking changes are likely

## Files

```
carp.go
cli.go
dependencies.go
```

## Carpfile

Dependencies are specified in a JSON file with the following format. This file is cumbersome to write directly, so I'd recommend creating a Python or SH executable file that echos the required JSON.

```
{
  // -- you can choose any group-names, but by default "main" is expected.
  "variables": {
    "requires": [
      /// -- list of variable dependencies
    ]
  },
  programs: {
    "requires": [
      // -- list of program dependencies
    ]
  },
  main: {
    requires: [
      // -- by default, the main group is checked. Depend on
      // -- other groups to organise dependencies tidily.
    ]
  }
}
```

## Dependencies

### "core/file"

A file dependency.

```json
{
  "type": "core/file",
  "path": "/home/alice/test.txt"
}
```

### "core/apt"

An apt-package dependency.

```json
{
  "type": "core/apt",
  "name": "git"
}
```

### "core/folder"

An folder dependency.

```json
{
  "type": "core/folder",
  "path": "/home/alice/code"
}
```

### "core/envvar"

An environment-variable dependency.

*Options:*
- name: the environmental variable name
- value: the expected value of the variable. Optional.

```json
{
  "type": "core/envvar",
  "name": "SHELL",
  "value": "/usr/bin/zsh
}
```

### "core/carpgroup"

Require all dependencies in another carp group to resolve.

```json
{
  "type": "core/carpgroup",
  "name": "my-variables"
}
```

### "core/snap"

A snap-package dependency.

```json
{
  "type": "core/snap",
  "name": "kale"
}
```

### "core/command"

A command dependency.

```json
{
  "type": "core/command",
  "name": "emoji"
}
```

### License

The MIT License

Copyright (c) 2020 Róisín Grannell

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
