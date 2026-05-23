# Exercise 1: Character Generator

## Goal

Generate one ASCII-art character from simple rules.

This exercise helps you understand that an ASCII-art character is just a group
of rows.

## Task

Complete:

```go
func GenerateBox(width int, height int) []string
```

## Rules

- Return an empty slice when width or height is less than `2`.
- The first and last row should be underscores.
- Middle rows should use pipes on the left and right.
- Every row must have the same width.

## Example

```go
GenerateBox(5, 4)
```

Expected:

```text
_____
|   |
|   |
_____
```

