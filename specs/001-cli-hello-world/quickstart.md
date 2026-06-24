# Quickstart Validation Guide: CLI Hello World

**Date**: 2026-06-23 | **Plan**: [plan.md](plan.md) | **Contract**: [contracts/cli-contract.md](contracts/cli-contract.md)

## Prerequisites

- Go 1.25+ installed (`go version` to verify)
- Repository cloned and current directory is the project root

## Build

```bash
go build -o hello-world ./cmd/hello-world
```

This produces the `hello-world` executable in the project root.

## Validation Scenarios

### 1. Happy Path — Greet by Name (FR-001, FR-002, FR-004)

```bash
./hello-world World
```

**Expected output** (stdout):
```
hello World
```

**Expected exit code**: `0`

**Verify exit code**:
```bash
./hello-world World; echo "exit: $?"
```

### 2. Missing Argument — Error Handling (FR-003)

```bash
./hello-world
```

**Expected output** (stderr):
```
Usage: hello-world <name>
```

**Expected exit code**: `1`

**Verify exit code**:
```bash
./hello-world 2>&1; echo "exit: $?"
```

### 3. Multi-Word Name — Quoted Argument (FR-005)

```bash
./hello-world "Jane Doe"
```

**Expected output** (stdout):
```
hello Jane Doe
```

### 4. Special Characters (FR-006)

```bash
./hello-world "O'Brien"
./hello-world "José"
```

**Expected output** (stdout):
```
hello O'Brien
hello José
```

### 5. Extra Arguments Ignored

```bash
./hello-world Alice Bob
```

**Expected output** (stdout):
```
hello Alice
```

## Running Tests

```bash
# Unit tests
go test ./internal/greeter/...

# Integration tests (builds and runs the binary)
go test ./tests/integration/...

# All tests with coverage
go test -cover ./...
```

## Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

Target: ≥80% line coverage on new code (Constitution II).
