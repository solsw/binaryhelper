# binaryhelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/binaryhelper.svg)](https://pkg.go.dev/github.com/solsw/binaryhelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/binaryhelper)

`binaryhelper.CopyFixed` copies between Go values with the same encoded size. Use it to read a little-endian byte sequence into a number or to copy fields between different struct types.

## Installation

Requires Go 1.26 or later.

```sh
go get github.com/solsw/binaryhelper
```

```go
import "github.com/solsw/binaryhelper"
```

For the GitLab module, use `github.com/solsw/binaryhelper` in both the install command and the import.

## API

### CopyFixed

```go
func CopyFixed(src, dst any) error
```

`CopyFixed` encodes `src` in little-endian order and decodes those bytes into `dst` using the same order. The encoded sizes must match.

| Argument | Accepted forms |
| --- | --- |
| `src` | A fixed-size value, a slice of fixed-size elements, or a pointer to either |
| `dst` | A pointer to a fixed-size value, a slice of fixed-size elements, or a pointer to such a slice |

A fixed-size value can be a boolean, a fixed-width numeric type such as `uint32` or `float64`, or an array or struct composed of fixed-size values. Struct fields are encoded in declaration order. Use a blank field (`_`) for explicit padding. Every other destination field must be exported, including fields in nested structs.

For slices, the encoded size is the current length multiplied by the encoded element size. A nil slice of fixed-size elements has encoded size zero.

#### Errors

`CopyFixed` returns an error for a nil interface or nil pointer, an unsupported source or destination, or unequal encoded sizes. On error, the destination retains its original value.

## Usage

### Copy between struct types

```go
package main

import (
	"fmt"

	"github.com/solsw/binaryhelper"
)

type Point struct{ X, Y int16 }
type Pair struct{ A, B int16 }

func main() {
	src := Point{X: 1, Y: 2}
	var dst Pair
	if err := binaryhelper.CopyFixed(src, &dst); err != nil {
		panic(err)
	}
	fmt.Println(dst) // {1 2}
}
```

### Decode a little-endian number

```go
var n int32
if err := binaryhelper.CopyFixed([]byte{0xd2, 0x04, 0x00, 0x00}, &n); err != nil {
	panic(err)
}
fmt.Println(n) // 1234
```

### Write into a slice

Pass a slice directly to fill its elements. The second call below uses a one-byte subslice of the same buffer, which is too short for a `uint16`.

```go
buf := make([]byte, 2)
if err := binaryhelper.CopyFixed(uint16(0x0102), buf); err != nil {
	panic(err)
}
fmt.Printf("% x\n", buf) // 02 01

err := binaryhelper.CopyFixed(uint16(0xabcd), buf[:1])
fmt.Printf("%t, % x\n", err != nil, buf) // true, 02 01
```

## See also

- [Package reference on pkg.go.dev](https://pkg.go.dev/github.com/solsw/binaryhelper)
- [encoding/binary](https://pkg.go.dev/encoding/binary)
