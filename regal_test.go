package regal

import (
	"testing"
)

func BenchmarkRegalEngine(b *testing.B) {

	var example2 = map[string]string{
		"ver1": "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5,10.1.1.6",
		"ver2": "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5",
		"ver3": "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5",
	}

	for i := 0; i < b.N; i++ {
		c2 := RegalEngine(
			example2,
			WithCombine(3),
			WithSchedule(2),
		)
		c2.Grouping()
	}
}
