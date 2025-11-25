<!-- 68aa0c09-3485-4567-ba93-4d2ded75dc28 8a999837-d0dd-4ec5-a53f-3a5f244f7988 -->
# Simplify GitHub Workflows

## Overview

Remove excessive automation and checks that create noise. Keep only essential CI (tests + build) and release workflows.

## Changes

### 1. Remove Unnecessary Workflows

Delete these workflow files:

- `.github/workflows/codeql.yml` - Security scanning (too defensive)
- `.github/workflows/stale.yml` - Auto-closes issues/PRs (creates noise)

### 2. Simplify CI Workflow

Modify `.github/workflows/ci.yml`:

- Remove the `golangci-lint` job entirely (lines 54-70)
- Keep only the `lint-and-test` job with:
- Go vet
- Go fmt check
- Tests with race detection
- Build verification

### 3. Remove Dependabot

Delete `.github/dependabot.yml` - currently configured to create up to 20 PRs per week for:

- Go modules (10 PRs)
- GitHub Actions (5 PRs)
- Docker images (5 PRs)

### 4. Keep Release Workflow

No changes to `.github/workflows/release.yml` - this is essential for:

- GoReleaser on version tags
- Docker image publishing

## Result

Minimal, focused CI/CD:

- **CI**: Basic Go checks (vet, fmt, test, build) on push/PR
- **Release**: Automated releases on git tags
- **No noise**: No auto-closing, no security scanning, no automated dependency PRs

### To-dos

- [ ] Delete .github/workflows/codeql.yml
- [ ] Delete .github/workflows/stale.yml
- [ ] Delete .github/dependabot.yml
- [ ] Remove golangci-lint job from .github/workflows/ci.yml