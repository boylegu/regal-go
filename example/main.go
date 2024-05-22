package main

import (
	"fmt"
	"github.com/boylegu/regal-go"
)

func main() {
	var example1 = [][]string{
		{"app-test-ver1", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
	}
	c1 := regal.RegalEngine(example1, regal.WithCombine(2))
	fmt.Println(c1.Grouping())

	var example2 = [][]string{
		{"ver1", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5,10.1.1.6"},
		{"ver2", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
		{"ver3", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
	}
	c2 := regal.RegalEngine(
		example2,
		regal.WithCombine(3),
		regal.WithSchedule(2),
		regal.WithPriorKey("ver2"),
	)
	for _, v := range c2.Grouping() {
		fmt.Println(v)
	}
}
