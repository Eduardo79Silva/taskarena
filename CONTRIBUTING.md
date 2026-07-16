# Contributing to TaskArena

Thanks for your interest in contributing! This document covers everything you need to know to get started.

## Ways to contribute

- Report bugs or suggest features by opening an [issue](https://github.com/Eduardo79Silva/taskarena/issues)
- Fix bugs or implement features via pull requests
- Improve documentation
- Share feedback on the scheduling algorithm, CLI ergonomics, or TUI design

## Getting started

1. **Fork** the repository and clone your fork:

   ```bash
   git clone https://github.com/<your-username>/taskarena.git
   cd taskarena
   ```

2. **Create a branch** for your change:

   ```bash
   git checkout -b feature/short-description
   # or: fix/short-description, docs/short-description
   ```

3. **Build and test locally**:
   ```bash
   go build ./...
   go vet ./...
   go test ./... -race -cover
   ```

## Development guidelines

- **Go version**: match whatever is declared in `go.mod`.
- **Formatting**: run `gofmt -w .` (or `goimports`) before committing. CI will flag unformatted code.
- **Package structure**: core logic lives in `internal/`, organized by responsibility (`task`, `scheduler`, `store`, `config`, `app`). The `app` package is the shared orchestration layer used by both the CLI (`internal/cli`) and the TUI (`tui/`) — new features should live there if the logic needs to be usable from both frontends.
- **Tests**: new logic should come with tests. Table-driven tests are preferred for anything with multiple input/output cases. Tests live alongside the package they test.
- **Commit messages**: write clear, descriptive messages. Conventional prefixes (`fix:`, `feat:`, `docs:`, `test:`, `refactor:`) are appreciated since they feed into the auto-generated changelog, but not strictly required.

## Submitting a pull request

1. Push your branch to your fork.
2. Open a PR against `main` in the upstream repo.
3. Fill out the PR template (what changed, why, how it was tested).
4. Ensure CI passes; PRs cannot merge until the `build-and-test` check is green.
5. Be responsive to review feedback. Conversations must be marked resolved before merge.

PRs are merged via squash merge, so your branch's individual commit history doesn't need to be pristine; focus on making the PR description clear instead.

## Reporting bugs

Please include:

- Your OS and `taskarena --version` output
- Steps to reproduce
- Expected vs. actual behavior
- Relevant config (`~/.config/taskarena/config.toml`), with any sensitive info removed

## Code of Conduct

This project follows a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you agree to uphold it.

## Questions

If anything here is unclear, feel free to open an issue asking — that's a sign the docs need improving.
