## hashbow

Somewhat a port of https://github.com/supercrabtree/hashbow to Go.
Uses https://github.com/lucasb-eyer/go-colorful underneath.

Documentation: http://godoc.org/github.com/fiatjaf/hashbow

## example

```go
package main

import (
    "github.com/fiatjaf/hashbow"
    "log"
)

func main () {
    log.Print(hashbow.Hashbow("Olivia Wilde"))
    log.Print(hashbow.Hashbow(22))
}
```

Will print `"#8b3c61"` and `"#3c61d8"`, which are beautiful colours.

## compatibility

This is not compatible with https://github.com/supercrabtree/hashbow in the sense that it will not generate the same colours from the same arguments. I want it to be compatible. Please submit a PR if know how to do it.
