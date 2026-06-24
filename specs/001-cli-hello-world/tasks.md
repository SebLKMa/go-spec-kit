# Tasks: CLI Hello World

**Input**: Design documents from `specs/001-cli-hello-world/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Tests are included — constitution requires unit tests (≥80% coverage) and integration tests for the primary user journey.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)
- Paths use Go project layout from plan.md: `cmd/hello-world/`, `internal/greeter/`, `tests/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Initialize Go module and project directory structure

- [x] T001 Initialize Go module with `go mod init` and create directory structure per plan.md (`cmd/hello-world/`, `internal/greeter/`, `tests/unit/`, `tests/integration/`)

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core greeting logic that all user stories depend on

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T002 Implement `Greet` function in `internal/greeter/greeter.go` that accepts a name string and returns the formatted greeting `hello <name>`

**Checkpoint**: Foundation ready — `Greet` function available for CLI entry point and tests

---

## Phase 3: User Story 1 — Greet a Named Person (Priority: P1) 🎯 MVP

**Goal**: User runs the program with a name argument and sees `hello <name>` on stdout

**Independent Test**: Build the binary, run `./hello-world World`, verify output is `hello World` and exit code is 0

### Tests for User Story 1

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [x] T003 [P] [US1] Unit test for `Greet` function in `tests/unit/greeter_test.go` — test with single name, multi-word name, and special characters (O'Brien, José)
- [x] T004 [P] [US1] Integration test for happy path in `tests/integration/cli_test.go` — build binary with `go build`, execute with a name argument, assert stdout matches `hello <name>` and exit code is 0

### Implementation for User Story 1

- [x] T005 [US1] Implement `main` function in `cmd/hello-world/main.go` — read `os.Args[1]`, call `greeter.Greet`, print result to stdout, exit 0
- [x] T006 [US1] Verify all US1 tests pass and greeting works for single names, multi-word quoted names, and special characters per contracts/cli-contract.md

**Checkpoint**: User Story 1 fully functional — `./hello-world Alice` outputs `hello Alice`

---

## Phase 4: User Story 2 — Handle Missing Argument (Priority: P2)

**Goal**: User runs the program with no arguments and sees an actionable usage message on stderr with exit code 1

**Independent Test**: Build the binary, run `./hello-world` with no args, verify stderr shows `Usage: hello-world <name>` and exit code is 1

### Tests for User Story 2

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [x] T007 [P] [US2] Integration test for missing argument in `tests/integration/cli_test.go` — execute binary with no arguments, assert stderr contains `Usage: hello-world <name>` and exit code is 1
- [x] T008 [P] [US2] Integration test for empty string argument in `tests/integration/cli_test.go` — execute binary with `""`, assert stderr contains usage message and exit code is 1

### Implementation for User Story 2

- [x] T009 [US2] Add argument validation to `cmd/hello-world/main.go` — check `len(os.Args) < 2` or empty first arg, print `Usage: hello-world <name>` to stderr, exit 1
- [x] T010 [US2] Verify all US2 tests pass and error handling works per contracts/cli-contract.md

**Checkpoint**: User Stories 1 AND 2 both work independently — happy path greets, missing arg shows usage

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Edge case handling, coverage verification, and quickstart validation

- [x] T011 [P] Integration test for extra arguments in `tests/integration/cli_test.go` — execute binary with `Alice Bob`, assert stdout is `hello Alice` (only first arg used)
- [x] T012 Run `go test -cover ./...` and verify ≥80% line coverage on new code per constitution
- [x] T013 Run `go vet ./...` and `gofmt -l .` to verify no linting or formatting violations
- [x] T014 Run quickstart.md validation scenarios end-to-end and verify all pass

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies — can start immediately
- **Foundational (Phase 2)**: Depends on Phase 1 — BLOCKS all user stories
- **User Story 1 (Phase 3)**: Depends on Phase 2 completion
- **User Story 2 (Phase 4)**: Depends on Phase 2 completion — can run in parallel with US1
- **Polish (Phase 5)**: Depends on Phases 3 and 4 being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) — no dependencies on other stories
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) — independent of US1 (different code path in main.go)

### Within Each User Story

- Tests MUST be written and FAIL before implementation
- Core logic before CLI wiring
- Implementation before verification

### Parallel Opportunities

- T003 and T004 (US1 tests) can run in parallel
- T007 and T008 (US2 tests) can run in parallel
- US1 and US2 phases can run in parallel after Foundational phase completes (they touch the same file but different code paths — if serialized, execute US1 first)

---

## Parallel Example: User Story 1

```bash
# Launch all tests for User Story 1 together:
Task: "Unit test for Greet function in tests/unit/greeter_test.go"
Task: "Integration test for happy path in tests/integration/cli_test.go"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (`go mod init`, directory structure)
2. Complete Phase 2: Foundational (`Greet` function)
3. Complete Phase 3: User Story 1 (tests → main.go → verify)
4. **STOP and VALIDATE**: `./hello-world World` → `hello World` ✓
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Working MVP!
3. Add User Story 2 → Test independently → Error handling complete
4. Polish → Coverage, linting, quickstart validation
5. Each story adds value without breaking previous stories

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story is independently completable and testable
- Verify tests fail before implementing
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
