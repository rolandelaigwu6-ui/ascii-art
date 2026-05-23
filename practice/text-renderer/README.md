# Exercise 6: Text Renderer

## Goal

Render text using a font map.

This is the full ASCII-art idea:

```text
for each row:
    for each character:
        add that character's row
```

## Task

Complete:

```go
func Render(input string, font map[byte][5]string) string
```

## Rules

- Return an empty string for empty input.
- Return an empty string for unknown characters.
- Support multiple input lines split by `\n`.
- Each rendered input line should produce 5 output rows.

## Example

Input:

```text
12
```

Expected:

```text
  |   ___ 
  |      |
  |   ___|
  |  |    
  |  |___ 
```

