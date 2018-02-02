# Valigator

## Overview 
validate standards for a project through specifications

A specification validates whether pattern is contained in a file and provides the error and resolution.

run `valigator.exe` with a `specifications.json` file at the root of your project:

``` json
[
    {
        "specification": "Use Angular version 5.2.0",
        "contains": true,
        "pattern": "\"@angular\/core\": \"\\^5\\.2\\.0\"",
        "filename": "package.json",
        "resolution": "Edit your package.json to dependecy for @angular/core to 5.2.0",
        "error": "@angular/core is not using version 5.2.0"
    },
    {
        "specification": "Never comment XML",
        "contains": false,
        "pattern": "<!--",
        "filename": "all",
        "resolution": "Remove XML comment from your file",
        "error": "File includes XML comment"
    }
]
```

## Compile

## Linux Bash 
1. `cd src`
2. `go build -o ../../bin/valigator.exe *.go`

## Windows Cmd
1. `cd src`
2. `go build && move valigator.exe ../../bin/`
