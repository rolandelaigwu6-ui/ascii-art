# Exercise 5: Font Loader

## Goal

Turn raw font rows into a usable font map.

In real ASCII-art projects, fonts usually come from files. This exercise avoids
file reading at first and focuses only on parsing rows.

## Task

Complete:

```go
func LoadFont(lines []string, chars []byte) map[byte][5]string
```

## Rules

- Each character uses exactly 5 lines.
- `chars[0]` gets the first 5 lines.
- `chars[1]` gets the next 5 lines.
- Return `nil` if the line count is not `len(chars) * 5`.

## Example

If `chars` is:

```go
[]byte{'1', '2'}
```

then lines `0` to `4` belong to `'1'`, and lines `5` to `9` belong to `'2'`.

