# gomodctl

_gomodctl_ - check and update go modules.

Currently supported commands:

- check - check project dependencies for the version information and shows outdated packages
- update - automatically sync project dependencies with their latest version

This is a fork of gomodctl which I removed the features I didn't use. I also rewrote some parts.

## Installation

Execute:

```bash
$ go get github.com/ormanli/gomodctl
```

Or using [Homebrew 🍺](https://brew.sh)

```bash
brew tap ormanli/gomodctl https://github.com/ormanli/gomodctl
brew install gomodctl
```

## Features

### gomodctl check

Check module versions in the given Go project.

Command:

```shell script
gomodctl check
```

Result:

```shell script
              MODULE              |       CURRENT       |       LATEST
----------------------------------+---------------------+----------------------
  github.com/stretchr/testify     | v1.3.0              | v1.4.0
  go.mongodb.org/mongo-driver     | v1.1.1              | v1.2.1
  github.com/mitchellh/go-homedir | v1.1.0              | v1.1.0
  github.com/ory/dockertest       | v3.3.5+incompatible | v3.3.5+incompatible
  github.com/pkg/errors           | v0.8.1              | v0.9.1
  github.com/spf13/cobra          | v0.0.5              | v0.0.5
  github.com/spf13/viper          | v1.4.0              | v1.6.2
----------------------------------+---------------------+----------------------
                                     NUMBER OF MODULES  |          7
                                  ----------------------+----------------------
```

Add `--json` parameter to the command to print result as a JSON.
Add `--path` parameter to the command to run command on another directory.

```shell script
gomodctl check --json --path ~/projects/gomodctl
```

### gomodctl update

Update module versions to latest minor

Command:

```shell script
gomodctl update
```

Result:

```shell script
Your dependencies updated to latest minor and go.mod.backup created
              MODULE              |      PREVIOUS       |         NOW
----------------------------------+---------------------+----------------------
  github.com/ory/dockertest       | v3.3.5+incompatible | v3.3.5+incompatible
  github.com/pkg/errors           | v0.8.1              | v0.9.1
  github.com/spf13/cobra          | v0.0.5              | v0.0.5
  github.com/spf13/viper          | v1.4.0              | v1.6.2
  github.com/stretchr/testify     | v1.3.0              | v1.4.0
  go.mongodb.org/mongo-driver     | v1.1.1              | v1.2.1
  github.com/mitchellh/go-homedir | v1.1.0              | v1.1.0
----------------------------------+---------------------+----------------------
                                     NUMBER OF MODULES  |          7
                                  ----------------------+----------------------
```

Add `--json` parameter to the command to print result as a JSON.
Add `--path` parameter to the command to run command on another directory.

```shell script
gomodctl update --json --path ~/projects/gomodctl
```

## How to ignore modules for version check and update

Create a `gomodctl.yaml` which has following structure which contains modules you want to ignore.
```yaml
ignored_modules:
 - github.com/x/y
 - github.com/a/b
```

gomodctl checks directories for `gomodctl.yaml` in given order.
 
1. `path` parameter
2.  current working directory
3.  home directory

## How to configure for private modules

Since check and update rely on go toolchain, if you have any private module that isn't publicly accessible, don't forget to set up your environment variables. For more information and how to configure, please check [Module configuration for non-public modules](https://golang.org/cmd/go/#hdr-Module_configuration_for_non_public_modules).
