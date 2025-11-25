<!-- 576439d9-d8f2-474c-845a-469eb58fd9c4 d4459855-9b9d-492c-b2e6-0ed0df09e6e6 -->
# Fix Go Version Mismatch

## Problem

The `go.mod` file specifies `go 1.24.7` (a development version set by your local Go 1.25.3), causing Docker builds to fail because CI/CD uses Go 1.22. We'll standardize on Go 1.23 (latest stable) for production/CI/CD. You can continue using Go 1.25 locally since Go is backward compatible.

## Files to Update

### 1. go.mod

Change line 3 from:

```go
go 1.24.7
```

to:

```go
go 1.23
```

### 2. Dockerfile  

Change line 5 from:

```dockerfile
FROM golang:1.22-alpine AS builder
```

to:

```dockerfile
FROM golang:1.23-alpine AS builder
```

### 3. .github/workflows/ci.yml

Change line 19 from:

```yaml
go-version: '1.22'
```

to:

```yaml
go-version: '1.23'
```

### 4. .github/workflows/release.yml

Change line 25 from:

```yaml
go-version: '1.22'
```

to:

```yaml
go-version: '1.23'
```

## Post-Update Steps

After making these changes:

1. Run `go mod tidy` to ensure dependencies are compatible
2. Verify `go.sum` is updated correctly
3. Test the build locally with `go build ./cmd/mcp-server-planton`
4. Commit and push to trigger CI/CD validation

### To-dos

- [ ] Update go.mod to use go 1.23
- [ ] Update Dockerfile to use golang:1.23-alpine
- [ ] Update CI workflow to use go-version 1.23
- [ ] Update Release workflow to use go-version 1.23
- [ ] Run go mod tidy to update dependencies
- [ ] Build and verify the application compiles successfully