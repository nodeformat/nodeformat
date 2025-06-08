# Golang Nodeformatter

This directory contains a nodeformatter for Golang code. It uses the standard `go/format` package to format Go source code.

## Usage

The primary function is `FormatNode(code string) (string, error)`.

### Example

```go
package main

import (
	"fmt"
	"log"

	"app/golang" // Assuming 'app' is your module name
)

func main() {
	unformattedCode := "package main; func  main (   ) { println(\"hello\") }" // Escaped quotes for Go string

	formattedCode, err := golang.FormatNode(unformattedCode)
	if err != nil {
		log.Fatalf("Error formatting code: %v", err)
	}

	fmt.Println(formattedCode)
}
```

This will output:

```go
package main

func main() { println("hello") }
```

## Error Handling

If the input code is not valid Go syntax, `FormatNode` will return an error.
```
