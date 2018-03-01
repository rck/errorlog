[![GO Report](https://goreportcard.com/badge/github.com/rck/errorlog)](https://goreportcard.com/report/github.com/rck/errorlog)
[![GoDoc](https://godoc.org/github.com/rck/errorlog?status.svg)](https://godoc.org/github.com/rck/errorlog)

# errorlog
This is a very simple package that collects `error`s. The advantage to a
simple slice where you add `error`s is that this one is concurrency safe,
while a slice/`append` is not.

# Installing
```
go get -u github.com/rck/errorlog
```
Next, include `errorlog` in your application:

```golang
import "github.com/rck/errorlog"
```

# Example

```golang
package main

import (
	"fmt"
	"sync"

	"github.com/rck/errorlog"
)

func main() {
	errs := errorlog.NewErrorLog()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			errs.Append(fmt.Errorf("Error number: %d", i))
		}(&wg, i)
	}
	wg.Wait()

	fmt.Println(errs.Len())
	for _, err := range errs.Errs() {
		fmt.Println(err)
	}
}
```
