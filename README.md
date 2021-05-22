# go-padding

`go-padding` generate a padded string according to the specified length.
We can specify whether we want string to be padded on the left or the right.

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
	padded, _ := padding.Pad(s, 20, "!", padding.PadRight)

	fmt.Printf("padded string: %s\n", padded)
	// padded string: hello!!!!!!!!!!!!!!!
}
```
