package regal

import (
	"strings"
	"testing"
)

func TestRegalEngine(t *testing.T) {
	combine := 2
	schedule := 3
	var example1 = [][]string{
		{"app-test-version1.0", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
	}

	if ans := RegalEngine(example1).Grouping(); len(ans) == 0 {
		t.Errorf("internal error")
	}

	if ans := RegalEngine(example1, WithCombine(combine), WithSchedule(schedule)).Grouping(); len(ans[0][1].([][]string)[1]) != combine {
		t.Errorf("combine expected be %d", combine)
	}

	if ans := RegalEngine(example1, WithCombine(combine), WithSchedule(schedule)).Grouping(); len(strings.Split(ans[0][1].([][]string)[0][0], ",")) != schedule {
		t.Errorf("schedule expected be %d", schedule)
	}

}

func BenchmarkRegalEngine(b *testing.B) {

	var example2 = [][]string{
		{"ver1", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5,10.1.1.6"},
		{"ver2", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
		{"ver3", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
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
