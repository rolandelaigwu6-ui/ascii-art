# Exercise 3: Font Builder

## Goal

Store several characters in a font map.

This exercise helps you understand this structure:

```go
map[byte][5]string
```

## Task

Complete:

```go
func BuildFont() map[byte][5]string
```

## Rules

- Return a map containing digits `0`, `1`, and `2`.
- Each digit must have exactly 5 rows.
- Each row must be exactly 5 characters wide.

## Expected Digit `0`

```text
 ___ 
|   |
|   |
|   |
|___|
```

## Expected Digit `1`

```text
  |  
  |  
  |  
  |  
  |  
```

## Expected Digit `2`

```text
 ___ 
    |
 ___|
|    
|___ 
```

