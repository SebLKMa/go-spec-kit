# CLI Contract: hello-world

**Date**: 2026-06-23 | **Plan**: [../plan.md](../plan.md)

## Command Schema

```
hello-world <name>
```

### Arguments

| Position | Name   | Required | Type   | Description                        |
|----------|--------|----------|--------|------------------------------------|
| 1        | name   | Yes      | string | The name to include in the greeting |

### Output

#### Success (exit code 0)

Writes to **stdout**:

```
hello <name>
```

- Output is a single line
- No trailing newline beyond the standard line terminator
- `<name>` is reproduced exactly as provided (no case transformation)

#### Error: Missing argument (exit code 1)

Writes to **stderr**:

```
Usage: hello-world <name>
```

### Exit Codes

| Code | Meaning                              |
|------|--------------------------------------|
| 0    | Greeting printed successfully        |
| 1    | Required argument missing            |

### Examples

```bash
# Happy path
$ hello-world World
hello World

# Named person
$ hello-world Alice
hello Alice

# Multi-word name (quoted)
$ hello-world "Jane Doe"
hello Jane Doe

# Missing argument
$ hello-world
Usage: hello-world <name>
# (exit code 1)

# Special characters
$ hello-world "O'Brien"
hello O'Brien
```
