# Changelog

All notable changes to this project will be documented in this file.

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
