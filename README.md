# go-example-app

Built with gin and gorm

# Run project

```
go run cmd/api/main.go
```

# Migrations

## sql-migrate

> SQL Schema migration tool for [Go](https://golang.org/). Based on [gorp](https://github.com/go-gorp/gorp) and [goose](https://bitbucket.org/liamstask/goose).

## Installation

To install the library and command line program, use the following:

```bash
go get -v github.com/rubenv/sql-migrate/...
```

For Go version from 1.18, use:

```bash
go install github.com/rubenv/sql-migrate/...@latest
```

## Usage

### As a standalone tool

```
$ sql-migrate --help
usage: sql-migrate [--version] [--help] <command> [<args>]

Available commands are:
    down      Undo a database migration
    new       Create a new migration
    redo      Reapply the last migration
    status    Show migration status
    up        Migrates the database to the most recent version available
```
