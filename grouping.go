package regal

import (
	"fmt"
	"sort"
	"strings"
)

type Converter interface {
	ConvertByOneSlice(i interface{}) ([]string, error)
	ConvertByDyadicSlice(i interface{}) ([][]string, error)
}

type InfaceToSliceConverter struct{}

func (c InfaceToSliceConverter) ConvertByOneSlice(i interface{}) ([]string, error) {
	switch v := i.(type) {
	case []string:
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", i)
	}
}

func (c InfaceToSliceConverter) ConvertByDyadicSlice(i interface{}) ([][]string, error) {
	switch v := i.(type) {
	case [][]string:
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", i)
	}
}

type regalList [][]interface{}

type baseInfo struct {
	versionHost *OrderedMap[string, string]
	params      paramOption
}

func (b *baseInfo) Grouping() regalList {
	vHost := b.initialize()
	return b.calculate(vHost)
}

func (b *baseInfo) initialize() regalList {
	var l regalList

	for el := b.versionHost.Front(); el != nil; el = el.Next() {
		ipList := strings.Split(el.Value, ",")
		l = append(l, []interface{}{el.Key, ipList})
	}
	return l
}

func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Pop(items []string) (string, []string) {
	i := len(items) - 1
	popItem := items[i]
	items[i] = popItem
	items[len(items)-1] = ""
	items = items[:len(items)-1]
	return popItem, items
}

func (b *baseInfo) setPrior(listPtr *regalList) (regalList, error) {
	if len(b.params.priorKey) > 0 {

		sort.SliceStable(*listPtr, func(i, j int) bool {
			return (*listPtr)[i][0] == b.params.priorKey
		})
	}
	return *listPtr, nil
}

func (b *baseInfo) calculate(vHost regalList) regalList {
	var c Converter = InfaceToSliceConverter{}

	baselist := make(regalList, len(vHost))
	baselistPtr := &baselist

	for hostindex := 0; hostindex < len(vHost); hostindex++ {
		convertToSlice, _ := c.ConvertByOneSlice(vHost[hostindex][1])
		hosts := convertToSlice[b.params.schedule:]
		Reverse(hosts)
		(*baselistPtr)[hostindex] = []interface{}{vHost[hostindex][0], [][]string{}}
		initHost := strings.Join(convertToSlice[:b.params.schedule], ", ")
		recursiveGrouping(hosts, &baselist, b, initHost, hostindex)
	}
	b.setPrior(baselistPtr)

	return *baselistPtr
}

func recursiveGrouping(hosts []string, baselist *regalList, b *baseInfo, init_host string, hostindex int) int {
	var grouping func(int) int
	var init_n int
	var popitem string
	var c Converter = InfaceToSliceConverter{}

	(*baselist)[hostindex][1] = [][]string{{init_host}}

	grouping = func(init_n int) int {
		f_count := init_n + 1

		if len(hosts) == 0 {
			return 0
		}

		convertToSlice, _ := c.ConvertByDyadicSlice((*baselist)[hostindex][1])
		(*baselist)[hostindex][1] = append(convertToSlice, []string{})
		for i := 1; i <= b.params.combine; i++ {
			defer func() {
				if r := recover(); r != nil {
				}
			}()
			popitem, hosts = Pop(hosts)
			convertToSlice, _ = c.ConvertByDyadicSlice((*baselist)[hostindex][1])
			convertToSlice[f_count] = append(convertToSlice[f_count], popitem)
		}
		return grouping(f_count)
	}
	return grouping(init_n)

}
