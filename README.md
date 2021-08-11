
# carp 🐟

carp is a simple dependency-checker that checks a host matches the expected configuration. It performs dependency-checks in parallel, so it's very fast.

## Usage

```zsh
carp --file <fpath>
```

## Stability Index

1, Experimental - This project might die, it's undertested and underdocumented, and redesigns and breaking changes are likely

## Build

To test carp, run:

```bash
go run *.go $HOME/carpfile.py --group main
```

To install carp, run:

```bash
cp carp /usr/bin/carp
```

## Files

```
carp.go
cli.go
dependencies.go
```

## Carpfile

Dependencies are specified in a JSON file with the format below. This file is cumbersome to write directly, so I'd recommend creating a Python or SH executable file that echos the required JSON.

Multiple machines can be configured by creating seperate carp groups; one for a laptop and one for a devbox in this case.

```
{
  "vars": {
    "requires": []
  },
  "gui": {
    "requires": [
      {
        "id": "core/file",
        "path": "/home/rg/.config/autostart/guake.desktop"
      }
    ]
  },
  "dotfiles": {
    "requires": [
      {
        "id": "core/file",
        "path": "/home/rg/.zshrc"
      }
    ]
  },
  "desktopFolders": {
    "requires": [
      {
        "id": "core/folder",
        "path": "/home/rg/Code"
      }
    ]
  },
  "files": {
    "requires": []
  },
  "aptRepos": {
    "requires": [
      {
        "id": "core/apt",
        "name": "build-essential"
      }
    ]
  },
  "uiAptRepos": {
    "requires": [
      {
        "id": "core/apt",
        "name": "ulauncher"
      }
    ]
  },
  "snapRepos": {
    "requires": [
      {
        "id": "core/snap",
        "name": "bashtop"
      }
    ]
  },
  "uiSnapRepos": {
    "requires": [
      {
        "id": "core/snap",
        "name": "chromium"
      }
    ]
  },
  "commands": {
    "requires": [
      {
        "id": "core/command",
        "name": "zoxide"
      }
    ]
  },
  "shared": {
    "requires": [
      {
        "id": "core/carpgroup",
        "name": "vars"
      },
      {
        "id": "core/carpgroup",
        "name": "commands"
      },
      {
        "id": "core/carpgroup",
        "name": "dotfiles"
      },
      {
        "id": "core/carpgroup",
        "name": "files"
      },
      {
        "id": "core/carpgroup",
        "name": "snapRepos"
      },
      {
        "id": "core/carpgroup",
        "name": "aptRepos"
      }
    ]
  },
  "laptop": {
    "requires": [
      {
        "id": "core/carpgroup",
        "name": "gui"
      },
      {
        "id": "core/carpgroup",
        "name": "shared"
      },
      {
        "id": "core/carpgroup",
        "name": "desktopFolders"
      },
      {
        "id": "core/carpgroup",
        "name": "uiSnapRepos"
      },
      {
        "id": "core/carpgroup",
        "name": "uiAptRepos"
      }
    ]
  },
  "devbox": {
    "requires": [
      {
        "id": "core/carpgroup",
        "name": "shared"
      }
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
