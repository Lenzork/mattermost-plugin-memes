# CI Configuration

This repository has two CI workflows:

## 1. Tests Workflow (`.github/workflows/test.yml`)

A comprehensive CI workflow that explicitly runs all tests and build verification:

- **Triggers**: Pull requests and pushes to `main`/`master` branches
- **Go Version**: 1.19 with module caching
- **Steps**:
  1. Download and verify Go module dependencies
  2. Build required tools (`build/manifest` and `build/deploy`)
  3. Run `golangci-lint` for code quality checks
  4. Execute all tests with race detection and coverage analysis
  5. Generate HTML coverage reports
  6. Build server binary for verification
  7. Create complete plugin bundle using `make dist`
  8. Upload coverage artifacts for download

### Test Coverage

The workflow runs tests for:
- `server/` - Main plugin server code (6 tests)
- `server/memelibrary/` - Meme library functionality (4 tests)

Coverage reports are generated and uploaded as artifacts for each CI run.

## 2. Community Plugin CI (`.github/workflows/ci.yml`)

Uses the shared Mattermost community plugin workflow that provides:
- Linting with golangci-lint
- Testing
- Building

## Local Development

To run the same checks locally:

```bash
# Run all tests
make test

# Run linting (requires golangci-lint installed)
make check-style

# Build and bundle plugin
make dist

# Generate coverage report
make coverage
```