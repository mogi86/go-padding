# go-padding

`go-padding` generate a padded string according to the specified length.
We can specify whether we want string to be padded on the left or the right.

For your reference, we can pad with `0` and spece by using [fmt](https://golang.org/pkg/fmt/) package.
Therefore, we don't need use `go-padding` if you just want to pad with `0` and spece.

## Installation

```bash
$ go get github.com/mogi86/go-padding
```

## Example Code

```go
package main

import (
	"fmt"
	"github.com/mogi86/go-padding"
)

func main() {
	s := "hello"
	padded, err := padding.Pad(s, 20, "!", padding.PadRight)
	if err != nil {
		fmt.Printf("padding failed: %v\n", err)
	}

	fmt.Printf("padded string: %s\n", padded)
}
```

- Output

```
padded string: hello!!!!!!!!!!!!!!!
```
