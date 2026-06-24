# Research: CLI Hello World

**Date**: 2026-06-23 | **Plan**: [plan.md](plan.md)

## Technology Decisions

### Go Standard Library for CLI Argument Parsing

- **Decision**: Use `os.Args` directly for argument parsing
- **Rationale**: The program requires exactly one positional argument. A flag-parsing library (e.g., `flag`, `cobra`, `urfave/cli`) adds unnecessary complexity for a tool with no flags or subcommands. `os.Args` is idiomatic Go for simple positional argument handling.
- **Alternatives considered**:
  - `flag` package — designed for named flags, not positional args; would require awkward `flag.Args()` after parsing no flags
  - `cobra` — full CLI framework; extreme overkill for a single-argument program
  - `urfave/cli` — lightweight but still adds a dependency for zero benefit here

### Project Layout

- **Decision**: `cmd/hello-world/main.go` + `internal/greeter/greeter.go`
- **Rationale**: Standard Go project layout. Separating the pure greeting function from the CLI entry point enables isolated unit testing of the greeting logic without invoking the binary. The `internal/` directory prevents external import of the greeter package (appropriate since this is a standalone tool, not a library).
- **Alternatives considered**:
  - Single `main.go` at root — simpler but conflates argument handling with business logic, making unit testing harder and violating single-responsibility (Constitution I)
  - `pkg/` instead of `internal/` — `pkg/` implies external consumption; not appropriate for a standalone CLI tool

### Testing Strategy

- **Decision**: Unit tests for the `greeter` package + integration tests that build and execute the compiled binary
- **Rationale**: Unit tests validate the greeting function in isolation. Integration tests validate the full user journey (compile → run → check output/exit code), satisfying Constitution II requirements for both unit and integration coverage. Contract tests are N/A (no public API or inter-service boundary).
- **Alternatives considered**:
  - Tests in `main` package only — cannot test greeting logic independently from CLI arg parsing
  - `TestMain` with `os.Exec` only — skips unit-level isolation

### Error Message Format

- **Decision**: Print `Usage: hello-world <name>` to stderr with exit code 1
- **Rationale**: Follows Unix conventions (usage messages to stderr, non-zero exit on error). Satisfies Constitution III (actionable error messages) and spec FR-003. The message tells the user exactly what went wrong and how to fix it.
- **Alternatives considered**:
  - Generic "error: missing argument" — less actionable, doesn't show correct usage
  - Print to stdout — violates Unix convention of separating normal output from error output

## Resolved Clarifications

No NEEDS CLARIFICATION items existed in the Technical Context. All technology choices were provided by the user (Go, compiled executable) or derived from the spec and constitution.
