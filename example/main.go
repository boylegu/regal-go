package main

import (
	"fmt"
	"regal"
)

func main() {
	var verHost = map[string]string{
		"app-test-version1.0": "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"}
	c := regal.RegalEngine(
		verHost,
		regal.WithCombine(2))
	//regal.WithSchedule(2))
	fmt.Println(c.Grouping())
}
