# prettypath
Pretty to print directory path, like shell `tree`.

# Installation
To install this package, you need to install Go and set your Go workspace first.

1. The first need Go installed (version 1.11+ is required), then you can use the below Go command to install fileserver.

```shell
$ go install github.com/jummyliu/prettypath@master
```

# Usage

```shell
$ prettypath -h

Usage of prettypath:
  -depth int
        max depth, -1 is unrestricted (default -1)
  -h    help
  -path string
        basic path (default ".")
```

```shell
$ prettypath -path .
.
|- Readme.md
|- go.mod
|- main.go

$ prettypath -path . -depth 3 > tree.txt

```
