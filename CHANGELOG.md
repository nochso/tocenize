# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this
project adheres to [Semantic Versioning](http://semver.org/).


## [Unreleased]

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


## 0.1.0 - 2016-11-06

### Added
- Initial public release.

[Unreleased]: https://github.com/nochso/tocenize/compare/0.1.0...HEAD