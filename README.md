### Golang playground

#### along with some notes

---
## Data types

 - `bool`: Boolean type representing true or false values.
- `string`: Represents a sequence of characters.
- `int`: Signed integer type, its size varies depending on the platform.
- `int8`: 8-bit signed integer.
- `int16`: 16-bit signed integer.
- `int32`: 32-bit signed integer.
- `int64`: 64-bit signed integer.
- `uint`: Unsigned integer type, its size varies depending on the platform.
- `uint8`: 8-bit unsigned integer (byte).
- `uint16`: 16-bit unsigned integer.
- `uint32`: 32-bit unsigned integer.
- `uint64`: 64-bit unsigned integer.
- `uintptr`: An unsigned integer type that is large enough to store the uninterpreted bits of a pointer value.
- `float32`: 32-bit floating-point number.
- `float64`: 64-bit floating-point number.
- `complex64`: Complex number with float32 real and imaginary parts.
- `complex128`: Complex number with float64 real and imaginary parts.
- `byte`: Synonym for `uint8`, often used in Go for ASCII characters.
- `rune`: Synonym for `int32`, represents a Unicode code point.
- `error`: Represents an error condition, commonly returned by functions indicating an error state.
- `pointer`: Represents the memory address of a variable.
- `array`: Fixed-size collection of elements of the same type.
- `slice`: Dynamically-sized, flexible view into elements of an array.
- `map`: Unordered collection of key-value pairs.
- `struct`: Composite data type consisting of fields with their own data types.
- `interface`: Defines a set of method signatures that a type must implement to satisfy the interface.
- `channel`: Concurrent communication primitive used to synchronize and communicate between goroutines.

---

Format specifiers
---
| Format Specifier | Data Type Represented |
|------------------|-----------------------|
| `%d`             | Signed integer (`int`, `int8`, `int16`, `int32`, `int64`) |
| `%b`             | Binary representation of an integer |
| `%o`             | Octal representation of an integer |
| `%x`, `%X`       | Hexadecimal representation of an integer |
| `%f`             | Floating-point number (`float32`, `float64`) |
| `%e`, `%E`       | Floating-point number in scientific notation |
| `%g`, `%G`       | Floating-point number in decimal or scientific notation |
| `%t`             | Boolean value (`true` or `false`) |
| `%c`             | Single Unicode character |
| `%s`             | String |
| `%q`             | Double-quoted string safely escaped with Go syntax |
| `%p`             | Pointer (displays the base 16 representation of the value) |

---
Working with Strings package
---

| Method      | Description                               | Signature                            |
|-------------|-------------------------------------------|--------------------------------------|
| Contains    | Reports whether substr is within s.       | `Contains(s, substr string) bool`    |
| Count       | Counts the number of non-overlapping instances of substr in s. | `Count(s, substr string) int`        |
| Index       | Returns the index of the first instance of substr in s, or -1 if substr is not present. | `Index(s, substr string) int`        |
| LastIndex   | Returns the index of the last instance of substr in s, or -1 if substr is not present. | `LastIndex(s, substr string) int`    |
| Replace     | Returns a copy of s with the first n non-overlapping instances of old replaced by new. | `Replace(s, old, new string, n int) string` |
| Split       | Splits s into substrings separated by sep and returns a slice of the substrings. | `Split(s, sep string) []string`      |
| ToLower     | Returns a copy of the string s with all Unicode letters mapped to their lower case. | `ToLower(s string) string`           |
| ToUpper     | Returns a copy of the string s with all Unicode letters mapped to their upper case. | `ToUpper(s string) string`           |
| Trim        | Returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed. | `Trim(s string, cutset string) string` |
| TrimLeft    | Returns a slice of the string s with all leading Unicode code points contained in cutset removed. | `TrimLeft(s string, cutset string) string` |
| TrimRight   | Returns a slice of the string s with all trailing Unicode code points contained in cutset removed. | `TrimRight(s string, cutset string) string` |
