# binaryhelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/binaryhelper.svg)](https://pkg.go.dev/github.com/solsw/binaryhelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/binaryhelper)

Package `binaryhelper` contains helpers for Go's [encoding/binary](https://pkg.go.dev/encoding/binary) package.

## Installation

```sh
go get github.com/solsw/binaryhelper
```

```go
import "github.com/solsw/binaryhelper"
```

## API

### CopyFixed

```go
func CopyFixed(src, dst any) error
```

`CopyFixed` reinterprets the binary representation of `src` into `dst`. The value of `src` is serialized with [binary.Write](https://pkg.go.dev/encoding/binary#Write) and the resulting bytes are deserialized into `dst` with [binary.Read](https://pkg.go.dev/encoding/binary#Read).

Both `src` and `dst` must be **fixed-size values** (or pointers to them) — booleans, numeric types other than `int`, `uint` and `uintptr`, and arrays or structs containing only such types — as required by `binary.Write` and `binary.Read`. Since the copy happens entirely in-process, byte order does not affect the result.

`dst` must be a pointer so that the result can be stored into it.

#### Errors

`CopyFixed` returns an error when:

- `src` or `dst` is `nil`;
- `src` is not a fixed-size value (reported by `binary.Write`);
- `dst` is not a pointer to a fixed-size value, or `dst` is larger than `src` (reported by `binary.Read`);
- `src` is larger than `dst` (bytes of `src` remain after filling `dst`).

On success the binary sizes of `src` and `dst` match exactly and `dst` holds the reinterpreted value.

## Usage

Copying between structs with identical binary layouts but different types:

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

Reinterpreting raw bytes as a number:

```go
var n int32
err := binaryhelper.CopyFixed([4]byte{0xD2, 0x04, 0x00, 0x00}, &n)
// n == 1234 (the bytes are interpreted as little-endian)
```

## See also

- [Package reference on pkg.go.dev](https://pkg.go.dev/github.com/solsw/binaryhelper)
- [encoding/binary](https://pkg.go.dev/encoding/binary)
