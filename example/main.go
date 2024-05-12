package main

import (
	"fmt"
	"regal"
)

func main() {
	//var example1 = [][]string{
	//	{"app-test-version1.0", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
	//}
	//c1 := regal.RegalEngine(
	//	example1,
	//	regal.WithCombine(2),
	//	regal.WithSchedule(3))
	//fmt.Println(c1.Grouping())

	var example2 = [][]string{
		{"ver1", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5,10.1.1.6"},
		{"ver2", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
		{"ver3", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
	}

	c2 := regal.RegalEngine(
		example2,
		regal.WithCombine(3),
		regal.WithSchedule(2),
	)
	for k, v := range c2.Grouping() {
		fmt.Println(k, v)
	}

}
