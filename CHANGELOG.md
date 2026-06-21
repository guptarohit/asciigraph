# Changelog

All notable changes to this project will be documented in this file.

## [0.11.0](https://github.com/guptarohit/asciigraph/compare/v0.10.0...v0.11.0) (2026-06-21)


### Added

* add ClearLines to redraw realtime graph without flicker ([66346cb](https://github.com/guptarohit/asciigraph/commit/66346cbe66650c72a056e8a4d91302f730a8e5e0))
* add spectrum/heatmap gradient coloring ([0b1d9ad](https://github.com/guptarohit/asciigraph/commit/0b1d9ad58b8f27234a9408ac06abd5f9b25a12ec))
* **cli:** add -xmin, -xmax, and -xt flags for X-axis support ([24e8e38](https://github.com/guptarohit/asciigraph/commit/24e8e38e9f8e1b6c509309091843a145a3983eb3))
* **cli:** clear only previous graph in realtime mode ([#79](https://github.com/guptarohit/asciigraph/issues/79)) ([3c935a3](https://github.com/guptarohit/asciigraph/commit/3c935a3d47b2f31d8615220c4c951632a6067f50))
* line ending option for raw terminals ([0148570](https://github.com/guptarohit/asciigraph/commit/0148570342114076f295e175c5bed332c00645a5))
* line ending option for raw terminals ([573072e](https://github.com/guptarohit/asciigraph/commit/573072ea42de2f4fac94a123f92530aa31100250))
* **options:** add X-axis configuration types and option functions ([4e5e0a8](https://github.com/guptarohit/asciigraph/commit/4e5e0a887cbd244532ca1512fa6abd58808c000f))
* **plot:** implement X-axis rendering with tick marks and labels ([6c1bb15](https://github.com/guptarohit/asciigraph/commit/6c1bb1549bf7a13c04bdaa7b515f81e1a9c43db6))


### Fixed

* 32-bit compile overflow and release workflow cleanup ([451bece](https://github.com/guptarohit/asciigraph/commit/451becef290c1c8e1d114d73cbe60a05fb5e7c0c))
* 32-bit compile overflow and release workflow cleanup ([#74](https://github.com/guptarohit/asciigraph/issues/74)) ([7f76123](https://github.com/guptarohit/asciigraph/commit/7f7612366a79f8ef909c872dcb5df677cf4a9fbb))
* **axis:** remove redundant maxWidth assignment ([134f68d](https://github.com/guptarohit/asciigraph/commit/134f68de5267b05048ddc682d0ab7e7a47dfbbbe))
* **cli:** clear previous graph before plotting in realtime mode ([cbc1f4a](https://github.com/guptarohit/asciigraph/commit/cbc1f4a0489cda336926b04b9676491e8a3d88d6)), closes [#44](https://github.com/guptarohit/asciigraph/issues/44)
* **color:** avoid Black sentinel output, guard NaN, copy gradient stops ([fa2c222](https://github.com/guptarohit/asciigraph/commit/fa2c22295f27cbe3ca7392490ec1df9b275a64ec))
* **color:** correct Black normalization and gray overflow in gradient ([989ac30](https://github.com/guptarohit/asciigraph/commit/989ac30e1a819cb0cd5cae6516f29a33a37e69ab))
* **color:** map bright grays to nearest of gray ramp or white ([f1278e6](https://github.com/guptarohit/asciigraph/commit/f1278e6dcae6917ab880caabaf69228a17cf3fd5))
* **color:** round gray ramp to nearest step in rgbToAnsi256 ([5d487f6](https://github.com/guptarohit/asciigraph/commit/5d487f68ffa8d687589c3c0c9e7572b5e1b0dfd4))
* display values for Y-axis when data points are all identical ([1e527b8](https://github.com/guptarohit/asciigraph/commit/1e527b808b1c8602a76e0b6554544ae1ac61f717))
* display values for Y-axis when data points are all identical ([41e8421](https://github.com/guptarohit/asciigraph/commit/41e8421e4a5062377a25a5b1de53924e451795f4)), closes [#61](https://github.com/guptarohit/asciigraph/issues/61)
* **docker:** add OCI labels to Dockerfile for ghcr.io metadata ([f83d907](https://github.com/guptarohit/asciigraph/commit/f83d9079b0aaac59b9bd15abe4a8cb26fcde9dbe))
* **docker:** add OCI labels to Dockerfile for ghcr.io metadata ([#76](https://github.com/guptarohit/asciigraph/issues/76)) ([02f3e4a](https://github.com/guptarohit/asciigraph/commit/02f3e4a0f5136d710be39c407f77cfffc384821b))
* leave legend boxes uncolored under gradient; clarify color docs ([24c6b2a](https://github.com/guptarohit/asciigraph/commit/24c6b2ade0e506bfb3c001724fd25a5afc2d139d))
* **legend:** prevent panic when only legends are set without series colors ([20e1ef8](https://github.com/guptarohit/asciigraph/commit/20e1ef88096cc36fe4739a0235deec7580499890))
* **legend:** prevent panic when only legends are set without series colors ([9220833](https://github.com/guptarohit/asciigraph/commit/92208336ddc6dab76d1e42a6a5a08e779472e762))
* **plot:** compute X-axis tick values in value-space, use %.2f when visible ticks have decimals ([aa3eff7](https://github.com/guptarohit/asciigraph/commit/aa3eff76e4a97b346689de23c83215384853576a))


### Changed

* **axis:** remove minNumLength sentinel from formatter path ([8a81a67](https://github.com/guptarohit/asciigraph/commit/8a81a6718d59db385822e454ecc2e5a387011845))
* **options:** add X-axis box-drawing characters to default CharSet ([6e65c25](https://github.com/guptarohit/asciigraph/commit/6e65c255366a63e04d98e1e573f34bd9f064eab3))
* **release:** update goreleaser config for v2 ([94f69c0](https://github.com/guptarohit/asciigraph/commit/94f69c07479ce38a32c6ef076757b77973d95197))
* Rename min/max variables to avoid builtin conflicts ([275d495](https://github.com/guptarohit/asciigraph/commit/275d495e7ab35e8f7f1c63759a9edc408b613bb2))

## [0.10.0](https://github.com/guptarohit/asciigraph/compare/v0.9.0...v0.10.0) (2026-06-21)


### Added

* threshold coloring: `ColorAbove`/`ColorBelow` options and `-ca`/`-cb` CLI flags ([#85](https://github.com/guptarohit/asciigraph/pull/85))
* spectrum/heatmap gradient coloring: `SeriesColorGradient` option, `HeatmapSpectrum` palette, and `-g` CLI flag ([#84](https://github.com/guptarohit/asciigraph/pull/84))
* flicker-free realtime mode via `ClearLines`, redrawing only the previous graph ([#79](https://github.com/guptarohit/asciigraph/pull/79)), closes [#44](https://github.com/guptarohit/asciigraph/issues/44)

## [0.9.0](https://github.com/guptarohit/asciigraph/compare/v0.8.1...v0.9.0) (2026-03-28)


### Added

* **plot:** add X-axis rendering with tick marks and labels ([#78](https://github.com/guptarohit/asciigraph/pull/78))
* **cli:** add -xmin, -xmax, and -xt flags for X-axis support ([#78](https://github.com/guptarohit/asciigraph/pull/78))


### Fixed

* **docker:** add OCI labels to Dockerfile for ghcr.io metadata ([#76](https://github.com/guptarohit/asciigraph/issues/76)) ([02f3e4a](https://github.com/guptarohit/asciigraph/commit/02f3e4a0f5136d710be39c407f77cfffc384821b))

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
