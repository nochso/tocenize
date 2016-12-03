Change Log
==========

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this
project adheres to [Semantic Versioning](http://semver.org/).

- [Unreleased](#unreleased)
- [0.3.0 - 2016-12-03](#030---2016-12-03)
- [0.2.0 - 2016-11-20](#020---2016-11-20)
- [0.1.0 - 2016-11-06](#010---2016-11-06)

<!--
Added      new features.
Changed    changes in existing functionality.
Deprecated once-stable features removed in upcoming releases.
Removed    deprecated features removed in this release.
Fixed      any bug fixes.
Security   invite users to upgrade in case of vulnerabilities.
-->

[Unreleased]
------------

### Fixed
- Inline `[link]`s are now stripped from TOC titles.


[0.3.0] - 2016-12-03
--------------------

### Added
- New flag `-indent` to change the indentation string for nested lists (default `\t`)

### Removed
- Removed verbosity flag `-v` and is now used for version display.

### Changed
- Flag `-V` is now `-v`, replacing verbosity.
- Diffs are now colourful and limited to 3 lines of context.


[0.2.0] - 2016-11-20
--------------------

### Added
- Added flag `-e` to update only existing TOCs.

### Changed
- `Document.Update()` now returns `Document, error` instead of just `error`
- Insert new line when inserting a TOC for the first time.
- Markdown is stripped from anchors and link texts.
    - Link texts get stripped of images and links (excl. text)
    - Anchors are stripped of all Markdown.
- Parse given file paths using [`filepath.Glob`](https://golang.org/pkg/path/filepath/#Glob).

### Fixed
- Fixed wrong behaviour of CRLF endings because of ineffectual assignment.
- Duplicate headings did not have unique anchor links. They are now numbered
  the same way as Github does it.

### Removed
- Removed flag `-u`. Updating is still used as a default when other flags are
  not set.


[0.1.0] - 2016-11-06
--------------------

### Added
- Initial public release.

[Unreleased]: https://github.com/nochso/tocenize/compare/0.3.0...HEAD
[0.3.0]: https://github.com/nochso/tocenize/compare/0.2.0...0.3.0
[0.2.0]: https://github.com/nochso/tocenize/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/nochso/tocenize/compare/37dbbf6741f917c976cc77cfc84be81ea5d86e7d...0.1.0