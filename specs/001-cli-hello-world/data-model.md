# Data Model: CLI Hello World

**Date**: 2026-06-23 | **Plan**: [plan.md](plan.md)

## Overview

This feature has no persistent data model. The program is stateless — it takes a string input, formats a greeting, and outputs it. No storage, database, or file I/O is involved.

## Entities

### Greeting

A transient value produced by combining a fixed prefix with a user-supplied name.

| Field  | Type   | Description                              | Validation                          |
|--------|--------|------------------------------------------|-------------------------------------|
| Name   | string | The name provided as a command line argument | Must be non-empty after trimming    |
| Output | string | The formatted greeting: `hello <Name>`   | Always in format `hello ` + Name    |

### State Transitions

None. The program has a single execution path:

1. **Start** → Read argument
2. **Argument present** → Format greeting → Print to stdout → Exit 0
3. **Argument missing** → Print usage to stderr → Exit 1
