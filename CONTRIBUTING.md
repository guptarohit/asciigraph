# Contributing

Thanks for contributing to `asciigraph`.

## Before opening a PR

- format Go code with `gofmt`
- run `go test ./...`
- keep changes focused and backward compatible where possible

## Pull request titles

This repository uses Conventional Commit style PR titles because release
automation relies on them.

Examples:

- `feat(axis): add YAxisValueFormatter`
- `fix(plot): preserve values for flat series`
- `docs: update README examples`
- `refactor: simplify legend padding`
- `test(plot): add regression coverage`

Recommended types:

- `feat`
- `fix`
- `docs`
- `test`
- `refactor`
- `perf`
- `build`
- `ci`
- `chore`
- `deps`
- `revert`

If a PR title is not in the expected format, maintainers may edit it before
merge.

## Commit messages

Conventional Commit style commit messages are welcome, but they are not
required.

Release automation relies on the pull request title, so contributors do not need
to rewrite individual commits just to match the release format.

Versioning notes:

- `feat` -> minor release
- `fix` -> patch release
- `!` or `BREAKING CHANGE` -> major release

## Merge guidance for maintainers

If merge commits are used, keep the final merge commit title aligned with the PR
title so release automation can infer the release correctly from git history.

## Release process

Releases are automated:

- `release-please` opens a release PR and updates `CHANGELOG.md`
- merging the release PR creates the tag and GitHub release
- GoReleaser publishes binaries, checksums, and container images

Please do not manually edit `CHANGELOG.md` for normal releases unless a
maintainer is intentionally correcting release notes.
