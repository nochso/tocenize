tocenize
========

[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/nochso/tocenize.svg)](https://github.com/nochso/tocenize/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nochso/tocenize)](https://goreportcard.com/report/github.com/nochso/tocenize)

Insert and maintain a table of contents in Markdown files.

- [Features](#features)
- [Installation](#installation)
	- [Pre-compiled binaries](#pre-compiled-binaries)
	- [From source](#from-source)
- [Usage](#usage)
- [Alternatives](#alternatives)
- [Changes](#changes)
- [License](#license)


Features
--------

tocenize generates a TOC (table of content) from Markdown files and inserts or
updates it in the given file.

- Cross-platform command line utility
  - Windows, Linux, Mac and *bsd (anything the Go compiler will handle)
- Anchor links are compatible to GFM (Github flavoured Markdown)
- Automatic "intelligent" insertion of new TOC
- Update existing TOCs without moving it
- Configurable max. and min. header depth
- Line endings are kept intact (LF or CRLF is detected and then used for new lines)


Installation
------------


### Pre-compiled binaries

Compiled binaries are available on the [releases page][releases].

Make sure to place it somewhere in your `$PATH`.


### From source

If you have a working Go environment, simply run:

```
go install github.com/nochso/tocenize/cmd/tocenize
```

If you've added `$GOPATH/bin` to your `$PATH`, you can now run `tocenize` from
anywhere.

Usage
-----

The output of `tocenize -h` should be self explanatory:

```
tocenize [options] FILE...

  -V    print version
  -d    print full diff to stdout
  -e    update only existing TOC (no insert)
  -max int
        maximum depth (default 99)
  -min int
        minimum depth (default 1)
  -p    print full result to stdout
  -v    verbose output
```


Alternatives
------------

- [github.com/axelbellec/gotoc](https://github.com/axelbellec/gotoc) inserts on
  top, doesn't update existing TOC, doesn't support setext-style headers


Changes
-------

All notable changes to this project will be documented in the [changelog].

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this
project adheres to [Semantic Versioning](http://semver.org/).


License
-------

This project is released under the [MIT license](LICENSE).


[changelog]: CHANGELOG.md
[releases]: https://github.com/nochso/tocenize/releases
[Go]: https://golang.org