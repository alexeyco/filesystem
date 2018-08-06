# File system abstraction

[![Travis](https://img.shields.io/travis/alexeyco/filesystem.svg)](https://travis-ci.org/alexeyco/filesystem)
[![Coverage Status](https://coveralls.io/repos/github/alexeyco/filesystem/badge.svg?branch=master)](https://coveralls.io/github/alexeyco/filesystem?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexeyco/filesystem)](https://goreportcard.com/report/github.com/alexeyco/filesystem)
[![GoDoc](https://godoc.org/github.com/alexeyco/filesystem?status.svg)](https://godoc.org/github.com/alexeyco/filesystem)
[![license](https://img.shields.io/github/license/alexeyco/filesystem.svg)](https://github.com/alexeyco/filesystem)

**Work in progress!**

1. [Install](#install)
1. [Filesystem manipulations](#filesystem-manipulations)
    1. [List directory contents](#list-directory-contents)
    1. [Check path](#check-path)
    1. [Path info](#path-info)
    1. [Create directories](#create-directories)
    1. [Rename](#move/rename)
    1. [Remove](#remove)
1. Seekers
1. [License](#license)

## Install

```
$ go get -u github.com/alexeyco/filesystem
```

## Filesystem manipulations
```go
root, err := filesystem.Root("/path/to/root/directory")
if err != nil {
    log.Fatalln(err)
}
```

### List directory contents
Directories:
```go
err := root.Each().Dir(func(dir *filesystem.Dir) {
    fmt.Println("Dir:  ", dir.Name())
})
```

Files:
```go
err = root.Each().File(func(file *filesystem.File) {
    fmt.Println("File: ", file.Name())
})
```

Anything:
```go
err = root.Each().Entry(func(entry filesystem.Entry) {
    fmt.Println("Entry: ", entry.Name())
})
```

### Check path
Check path exists:
```go
exist := root.Exist("path/to/file")
```

Check path is a directory:
```go
exist := root.IsDir("path/to/directory")
```

Check path is a file:
```go
exist := root.IsFile("path/to/file")
```

### Path info
Get directory:
```go
dir, err := root.Dir("path/to/dir")
```

Get file:
```go
file, err := root.Dir("path/to/file")
```

### Create directories
Create directory:
```go
err := root.Mkdir("path/to/directory")
```

### Move/rename
Rename files and directories:
```go
err := root.Move("path/to/source", "path/to/dest")
```

### Remove
Remove files and directories:
```go
err := root.Move("path/to/source", "path/to/dest")
```

## License

```
MIT License

Copyright (c) 2018 Alexey Popov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
