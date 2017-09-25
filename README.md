# gregexp

[![Build Status](https://travis-ci.org/rosylilly/gregexp.svg?branch=master)](https://travis-ci.org/rosylilly/gregexp)
[![Coverage Status](https://coveralls.io/repos/github/rosylilly/gregexp/badge.svg?branch=master)](https://coveralls.io/github/rosylilly/gregexp?branch=master)
[![Godoc](https://godoc.org/github.com/rosylilly/gregexp?status.svg)](https://godoc.org/github.com/rosylilly/gregexp)

`gregexp` provides convertor for Shell file grep pattern to Regexp.

## Usage

```go
package main

import (
	"github.com/rosylilly/gregexp"
	"fmt"
)

func main() {
	reg, err := gregexp.Convert("lib/**/*.go")
	if err != nil {
		fmt.Error(err)
	}

	if reg.MatchString("lib/gregexp.go") {
		fmt.Println("matched")
	} else {
		fmt.Println("not matched")
	}
}
```

## License

[MIT License](https://github.com/rosylilly/gregexp/blob/master/LICENSE)
