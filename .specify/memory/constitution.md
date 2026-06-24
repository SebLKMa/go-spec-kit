<!--
SYNC IMPACT REPORT
==================
Version change:        0.0.0 (template) → 1.0.0
Modified principles:   [template placeholders] → 4 concrete principles (template had 5 slots; 4 used per user request)
Added sections:        Core Principles (I–IV), Quality Gates, Development Workflow, Governance
Removed sections:      None (template placeholder slots reduced from 5 to 4)
Templates requiring updates:
  .specify/templates/plan-template.md    ✅ aligned — Constitution Check section dynamically references principles
  .specify/templates/spec-template.md    ✅ aligned — Success Criteria maps to Performance / UX principles
  .specify/templates/tasks-template.md   ✅ aligned — Phase/quality-gate structure mirrors Quality Gates section
Follow-up TODOs:       None — all placeholders resolved.
-->

# hello-spec-it Constitution

## Core Principles

### I. Code Quality

All code MUST be clean, readable, and maintainable. Every function and module MUST have a single,
clear responsibility. Pull requests require at least one approval before merging — no exceptions.
Technical debt MUST be tracked and addressed within the sprint it is introduced. Dead code MUST be
deleted, not commented out. Naming MUST be descriptive and self-documenting; abbreviations are
forbidden unless universally understood in the domain.

**Rationale**: Readable code reduces onboarding friction, lowers defect rates, and enables
confident refactoring as the system evolves.

### II. Testing Standards

Tests MUST be written before or alongside implementation — never as an afterthought. Every feature
MUST include:

- Unit tests with ≥80% line coverage on new code
- At least one integration test covering the primary user journey
- Contract tests for any public API or inter-service boundary

Tests MUST be deterministic and isolated. Flaky tests MUST be fixed or removed within one sprint of
detection. Test names MUST describe the scenario in plain language: what is given, what action is
taken, and what outcome is expected.

**Rationale**: Testable code is better-designed code. Automated tests are the only reliable
regression guard at scale and make refactoring safe.

### III. User Experience Consistency

All user-facing interfaces MUST conform to a single, documented design system. Specifically:

- Error messages MUST be actionable: they MUST state what went wrong and how to recover.
- Navigation patterns, labeling conventions, and interaction models MUST be consistent across all
  features.
- New UI patterns require approval before introduction; ad-hoc patterns are forbidden.
- Breaking changes to established UX patterns MUST be explicitly signed off and communicated to users.

**Rationale**: Consistent UX reduces cognitive load, decreases support burden, and builds the user
trust that drives retention.

### IV. Performance Requirements

Every feature MUST define measurable performance targets in its spec before implementation begins.
Baseline requirements that apply unless a spec overrides them with documented justification:

- API endpoints MUST respond within **200ms at p95** under expected load.
- UI interactions MUST NOT block the main thread for more than **100ms**.
- Page / screen initial load MUST complete within **2 seconds** on a standard connection.
- Performance regressions of **>10% versus the established baseline** MUST be justified or
  reverted before the feature ships.

**Rationale**: Performance is a feature. Unmeasured performance is unmanaged performance, and
regressions compound silently until they become crises.

## Quality Gates

Every pull request MUST satisfy all of the following before merge:

- All tests pass (unit, integration, contract).
- No new linting or static-analysis violations introduced.
- Code coverage on new code does not fall below the 80% threshold.
- Performance benchmarks have been run and deltas reviewed against baseline.
- UX changes have been reviewed against the design system.
- No `TODO` / `FIXME` comments exist without a linked, open issue.
- The change has at least one approving review from a team member.

## Development Workflow

1. Feature work begins with a spec (`/speckit-specify`) and an implementation plan
   (`/speckit-plan`).
2. Tasks are generated from the plan (`/speckit-tasks`) and executed in priority order.
3. Each user story MUST be independently testable and demo-able before work on the next begins.
4. Commits MUST be atomic: one logical change per commit with a descriptive message in imperative
   mood (e.g., "Add rate limiter to auth endpoint").
5. The `main` branch MUST always be in a deployable state; no work-in-progress merges.

## Governance

This constitution supersedes all other project practices and conventions. Amendments require:

1. A written proposal describing the change and its rationale.
2. Review and approval from at least two team members.
3. A migration plan for any existing code that violates the amended principle.
4. A version bump per the policy below.

**Versioning policy**:

- **MAJOR**: Backward-incompatible changes — principle removals or redefinitions.
- **MINOR**: New principle or section added, or materially expanded guidance.
- **PATCH**: Clarifications, rewording, or typo fixes with no semantic change.

All PRs and code reviews MUST verify compliance with this constitution. Complexity violations MUST
be justified in the plan's Complexity Tracking table before work begins.

**Version**: 1.0.0 | **Ratified**: 2026-06-23 | **Last Amended**: 2026-06-23
