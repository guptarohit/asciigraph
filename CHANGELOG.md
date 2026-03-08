# Changelog

All notable changes to this project will be documented in this file.

## [0.8.1](https://github.com/guptarohit/asciigraph/compare/v0.8.0...v0.8.1) (2026-03-08)


### Fixed

* 32-bit compile overflow and release workflow cleanup ([451bece](https://github.com/guptarohit/asciigraph/commit/451becef290c1c8e1d114d73cbe60a05fb5e7c0c))
* 32-bit compile overflow and release workflow cleanup ([#74](https://github.com/guptarohit/asciigraph/issues/74)) ([7f76123](https://github.com/guptarohit/asciigraph/commit/7f7612366a79f8ef909c872dcb5df677cf4a9fbb))


### Changed

* **axis:** remove minNumLength sentinel from formatter path ([8a81a67](https://github.com/guptarohit/asciigraph/commit/8a81a6718d59db385822e454ecc2e5a387011845))

## [0.8.0] - (2026-03-08)


### Added

* Option to format Y-axis values with `YAxisValueFormatter(...)` ([#58](https://github.com/guptarohit/asciigraph/pull/58))
* Option to customize plot characters with `SeriesChars(...)` and `CreateCharSet(...)` ([#70](https://github.com/guptarohit/asciigraph/pull/70))
* CLI: `-x` flag to specify custom characters, including comma-separated values for multiple series ([#70](https://github.com/guptarohit/asciigraph/pull/70))
* Option to configure line endings for raw terminals ([#71](https://github.com/guptarohit/asciigraph/pull/71))


### Fixed

* Respect caller-specified precision for large numbers ([#69](https://github.com/guptarohit/asciigraph/pull/69))
* Preserve exact Y-axis values for series with identical data points ([#65](https://github.com/guptarohit/asciigraph/pull/65)), closes [#61](https://github.com/guptarohit/asciigraph/issues/61)
* Prevent panics when legends are set without series colors ([#64](https://github.com/guptarohit/asciigraph/pull/64))

### Changed

* Automates the release flow, update goreleaser config for v2 ([#72](https://github.com/guptarohit/asciigraph/pull/72))

## [0.7.3] - 2024-10-26

### Fixed

- Incorrect plot height calculation for small value ranges (#59)

## [0.7.2] - 2024-08-12

### Fixed

- Unintended modification of input data (#55)

## [0.7.1] - 2024-03-30

### Added

- CLI: Option to specify legends for series (`sl`)

## [0.7.0] - 2024-03-30

### Added

- CLI: Options to specify delimiter (`d`) and number of series (`sn`)

### Changed

- CLI: Option (`sc`) to specify series colors

## [0.6.0] - 2024-03-25

### Added

- Option to add legends for colored graphs

## [0.5.6] - 2023-06-24

### Added

- Options to set upper & lower bound of graph

## [0.5.5] - 2022-05-03

### Added

- Ansi colors support for graphs

## [0.5.4] - 2022-05-03

### Added

- Option to plot multiple series together (#34)
- Dockerfile file support (#33)

## [0.5.3] - 2022-02-20

### Fixed

- Handled NaN first value (#32)
- Fixed incorrect y-axis start value tick (#31)

## [0.5.2] - 2021-03-28

### Added

- added support to set custom precision of data point labels along the y-axis
- added go module support

### Changed

- updated README to markdown format

## [0.5.1] - 2020-09-14

### Added

- added support for NaN values in series
- added option to control fps of plot rendering via cli for real-time data

### Changed

- removed use of append() method
- make caption centered
- removed trailing spaces from plot

## [0.5.0] - 2020-06-28

### Added

- added support for the realtime plot of data points (from stdin) for CLI.

## [0.4.2] - 2020-06-07

### Fixed

- Prevent panics when data is flat. (#8)
- Prevent BADPREC issue when maximum and minimum values in a series are 0. (#10)

[0.8.0]: https://github.com/guptarohit/asciigraph/releases/tag/v0.8.0
[0.7.3]: https://github.com/guptarohit/asciigraph/releases/tag/v0.7.3
[0.7.2]: https://github.com/guptarohit/asciigraph/releases/tag/v0.7.2
[0.7.1]: https://github.com/guptarohit/asciigraph/releases/tag/v0.7.1
[0.7.0]: https://github.com/guptarohit/asciigraph/releases/tag/v0.7.0
[0.6.0]: https://github.com/guptarohit/asciigraph/releases/tag/v0.6.0
[0.5.6]: https://github.com/guptarohit/asciigraph/releases/tag/v0.5.6
[0.5.5]: https://github.com/guptarohit/asciigraph/releases/tag/v0.5.5
[0.5.4]: https://github.com/guptarohit/asciigraph/releases/tag/v0.5.4
[0.5.3]: https://github.com/guptarohit/asciigraph/releases/tag/v0.5.3
[0.5.2]: https://github.com/guptarohit/asciigraph/releases/tag/v0.5.2
[0.5.1]: https://github.com/guptarohit/asciigraph/releases/tag/v0.5.1
[0.5.0]: https://github.com/guptarohit/asciigraph/releases/tag/v0.5.0
[0.4.2]: https://github.com/guptarohit/asciigraph/releases/tag/v0.4.2
