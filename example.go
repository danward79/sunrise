package main

import (
	"fmt"
	"github.com/danward79/sunrise"
)

func main() {
	//Melbourne -37.81, 144.96
	melbourne := sunrise.NewLocation(-37.81, 144.96)

	formatStr := "Jan 2 15:04:05"

	melbourne.Today()
	fmt.Println(melbourne.Sunrise().Format(formatStr))
	fmt.Println(melbourne.Sunset().Format(formatStr))
}
