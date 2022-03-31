## Command line tool to manage computer files
`brt` is a command line file management tool that provides simple and verbose commands

Install the package to get started\
`go install github.com/wariored/brt/cmd/brt`

## Usage
To add a file:\
`brt af filename`

You can specify the path or subpath\
`brt af path/to/the/file/filename`


To copy a content of a file to a new file\
`brt cf source_filename target_filename`

To get information about existing commands type `brt`
```
Description:
    File management tool

Sub-commands:
    brt add-file      Add a file, shortcut: af
    brt add-dir       Add a directory , shortcut: ad
    brt copy-file     Copy a file, shortcut: cf
    brt copy-dir      Copy a directory, shortcut: cd
    brt remove-dir    Remove a directory, shortcut: rd
    brt remove-file   Remove a file, shortcut: rf
```
