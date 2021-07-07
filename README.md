# replace

Command line tool for search and replace of text.

Note: Files must be checked into git for this to work on them. The use-case for the tool is to edit source code, e.g. batch renaming of variables.

# Installation

```sh
go get github.com/joakimofv/replace
```

# Syntax

First arg: old pattern
Second arg: new pattern
Third+ args: filepaths or filepath patterns

'*' is interpreted as a wildcard. For the filepath pattern '**' can also be used to represent any number of directories.

# Usage examples

```sh
replace 'MyFunction(*)' 'MyFunction(*, true)'
replace 'MyFunction(*)' 'MyFunction(*, true)' file1.go file2.go pkg/*
```
