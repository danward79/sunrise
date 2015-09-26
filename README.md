sunrise
=======

This is a modified version of the code from [keep94](github.com/keep94/sunrise) I tweaked it for my application.

Calculates sunrise and sunset times

This API is now stable.

Using
-----

```
import "github.com/danward79/sunrise"
```

Installing
----------

```
go get github.com/danward79/sunrise
```

Online Documentation
--------------------

Online documentation is available [here](http://godoc.org/github.com/danward79/sunrise).

##Example below

```
package main

import (
	"fmt"
	"github.com/danward79/sunrise"
)

func main() {
	//Melbourne -37.81, 144.96
	melbourne := sunrise.NewLocation(-37.81, 144.96)
	fmt.Println(melbourne)

	formatStr := "Jan 2 15:04:05"

	melbourne.Today()
	fmt.Println(melbourne.Sunrise().Format(formatStr))
	fmt.Println(melbourne.Sunset().Format(formatStr))
}
```
