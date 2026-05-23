# Exercise 4: Input Validator

## Goal

Check whether a string can be converted before rendering it.

## Task

Complete:

```go
func IsValidInput(input string, font map[byte][5]string) bool
```

## Rules

- Empty input is invalid.
- Newline `\n` is allowed.
- Empty lines are invalid.
- Every non-newline character must exist in the font map.

## Examples

```go
IsValidInput("12", font)    // true
IsValidInput("1\n2", font)  // true
IsValidInput("1a", font)    // false
IsValidInput("", font)      // false
IsValidInput("1\n", font)   // false
```

