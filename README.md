# compose-to-intellij

[![Build](https://github.com/J-R-Oliver/compose-to-intellij/actions/workflows/build.yml/badge.svg)](https://github.com/J-R-Oliver/compose-to-intellij/actions/workflows/build.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/J-R-Oliver/compose-to-intellij)](https://github.com/gomods/athens)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/J-R-Oliver/compose-to-intellij)](https://github.com/J-R-Oliver/compose-to-intellij/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/J-R-Oliver/compose-to-intellij.svg)](https://pkg.go.dev/github.com/J-R-Oliver/compose-to-intellij)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](http://unlicense.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/J-R-Oliver/compose-to-intellij)](https://goreportcard.com/report/github.com/J-R-Oliver/compose-to-intellij)

<table>
<tr>
<td>
Command line application to convert Docker Compose environment variables to IntelliJ format.
</td>
</tr>
</table>

## Intro

`compose-to-intellij` converts environment variables configured in [compose files](https://compose-spec.io) to the string 
format used by iDEA products such as `IntelliJ` or `WebStorm`. Compose environment variables are set using the 
`environment` key as part of `service` declarations. `compose-to-intellij` will output an `IntelliJ` environment 
variable string for each `service` configured in the`compose` file.  Optional arguments of compose service names can be 
passed to filter the output.

## Contents

- [Installation](#installation)
- [Arguments](#arguments)
- [Options](#options)
- [Local Development](#local-development)
- [Testing](#testing)
- [Conventional Commits](#conventional-commits)
- [GitHub Actions](#github-actions)

## Installation

`compose-to-intellij` can be installed either with [Homebrew](https://brew.sh), built from source, or downloaded from 
GitHub as a [GitHub release asset](https://github.com/J-R-Oliver/compose-to-intellij/releases).

### Homebrew

`compose-to-intellij` can be installed with `Homebrew` either by tapping [j-r-oliver/tools](https://github.com/J-R-Oliver/homebrew-tools) 
and then installing the `formulae`...

```shell
brew tap j-r-oliver/tools
brew install compose-to-intellij
```

...or installing directly from the tap.

```shell
brew install j-r-oliver/tools/compose-to-intellij
```

### From Source

To start `clone` the repository to your local machine. The following command will build a native executable and output
it to `/dist`:

```shell
go build -o dist/compose-to-intellij cmd/compose-to-intellij/main.go
```

## Arguments

Optional arguments can be passed after any options to filter the output. These arguments should match the `service` 
keys defined in the `compose` file. For Example:

```shell
compose-to-intellij application database
```

## Options

`compose-to-intellij` has a handful of options. These options can be used to override the defaults that have been 
provided. For example:

```shell
compose-to-intellij -i ./build/docker-compose.build.yaml
```

### Command Line Options

The following command line options are available for configuration:

| Option               | Default                         | Description                           |
|----------------------|---------------------------------|---------------------------------------|
| -i, --input \<input> | ./openapi-validator-report.json | filepath for docker-compose YAML file |
| -v, --version        |                                 | output the version number             |
| -h, --help           |                                 | display help for command              |

## Local Development

### Prerequisites

To install and modify this project you will need to have:

- [Go](https://go.dev)
- [Git](https://git-scm.com)

### Installation

To start, please `fork` and `clone` the repository to your local machine.

## Testing

All tests have been written using the [testing](https://pkg.go.dev/testing) package from the
[Standard library](https://pkg.go.dev/std). To run the tests execute:

```shell
go test -v ./...
```

Code coverage is also measured by using the `testing` package. To run tests with coverage execute:

```shell
go test -coverprofile=coverage.out  ./...
```

## Conventional Commits

This project uses the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification for commit
messages. The specification provides a simple rule set for creating commit messages, documenting features, fixes, and
breaking changes in commit messages.

A [pre-commit](https://pre-commit.com) [configuration file](.pre-commit-config.yaml) has been provided to automate
commit linting. Ensure that *pre-commit* has been [installed](https://www.conventionalcommits.org/en/v1.0.0/) and
execute...

```shell
pre-commit install
````

...to add a commit [Git hook](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks) to your local machine.

An automated pipeline job has been [configured](.github/workflows/build.yml) to lint commit messages on a push.

## GitHub Actions

A CI/CD pipeline has been created using [GitHub Actions](https://github.com/features/actions) to automated tasks such as
linting and testing.

### Build Workflow

The [build](./.github/workflows/build.yml) workflow handles integration tasks. This workflow consists of two jobs, `Git`
and `Go`, that run in parallel. This workflow is triggered on a push to a branch.

#### Git

This job automates tasks relating to repository linting and enforcing best practices.

#### Go

This job automates `Go` specific tasks.
