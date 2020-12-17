#!/usr/bin/env python3

import json
import sys
import os
from pathlib import Path

home = str(Path.home())

def depend_file(path):
  return {
    "id": "core/file",
    "path": path
  }

def depend_apt(name):
  return {
    "id": "core/apt",
    "name": name
  }

def depend_folder(path):
  return {
    "id": "core/folder",
    "path": path
  }

def depend_envvar(name, value):
  return {
    "id": "core/envvar",
    "name": "SHELL",
    "value": value
  }

def depend_carpgroup(name):
  return {
    "id": "core/carpgroup",
    "name": name
  }

def depend_snap(name):
  return {
    "id": "core/snap",
    "name": name
  }

def depend_command(name):
  return {
      "id": "core/command",
      "name": name
  }

def list_dependencies():
  return {
      "vars": {
        "requires": [
           depend_envvar("SHELL", "/usr/bin/zsh"),
           depend_envvar("EDITOR", "vim")
        ]
      },
      "dotfiles": {
        "requires": [
          depend_file(os.path.join(home, "/.zshrc")),
          depend_file(os.path.join(home, "/.zshrc.greetings")),
          depend_file(os.path.join(home, "/.zshrc.functions")),
          depend_file("/usr/bin/yadm")
        ]
      },
      "homeFolders": {
        "requires": [
          depend_file(os.path.join(home, "Code")),
          depend_file(os.path.join(home, "Drive"))
        ]
      },
      "aptRepos": {
        "requires": [
          depend_apt("asciinema"),
          depend_apt("build-essential"),
          depend_apt("catimg"),
          depend_apt("fzf"),
          depend_apt("git-lfs"),
          depend_apt("guake"),
          depend_apt("imagemagick"),
          depend_apt("python-pip"),
          depend_apt("python-pip3")
        ]
      },
      "npmRepos": {
        "requires": [
          depend_command("emoji")
        ]
      },
      "snapRepos": {
        "requires": [
          depend_snap("audacity"),
          depend_snap("chromium"),
          depend_snap("code"),
          depend_snap("discord"),
          depend_snap("gimp"),
          depend_snap("hotline"),
          depend_snap("htop"),
          depend_snap("kale"),
          depend_snap("kohl"),
          depend_snap("polonium"),
          depend_snap("rpgen"),
          depend_snap("snap-store"),
          depend_snap("snapcraft"),
          depend_snap("vlc")
        ]
      },
      "main": {
        "requires": [
          depend_carpgroup("vars"),
          depend_carpgroup("dotfiles"),
          depend_carpgroup("homeFolders"),
          depend_carpgroup("snapRepos")
        ]
      }
  }

def main():
  deps = list_dependencies()
  print(json.dumps(deps, indent=2))

main()
