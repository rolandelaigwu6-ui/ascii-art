# String Art Converter Practice

## Goal

Build a function that converts digit strings into ASCII-art.

You are practicing this idea:

```text
To print digits side by side, combine row 0 of every digit,
then row 1 of every digit, then row 2, and so on.
```

## Task

Complete this function:

```go
func StringToArt(input string) string
```

## Rules

- Support digits only.
- Start by supporting `1`, `2`, and `3`.
- Each digit must be 5 characters wide.
- Each digit must be 5 rows tall.
- Digits should be joined without extra spacing.
- Return an empty string for invalid input.

## Expected Example

Input:

```text
13
```

Output:

```text
  |   ___ 
  |      |
  |   ___|
  |      |
  |   ___|
```

## Reasoning Questions

Before coding, answer these in your notes:

1. Why is each digit stored as 5 strings?
2. Why do we loop over rows first?
3. Why do we loop over the input characters inside the row loop?
4. What should happen if the input contains `a`?
5. What should happen if the input is empty?
