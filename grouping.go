package regal

import (
	"strings"
)

type regalList [][]interface{}

type baseInfo struct {
	versionHost map[string]string
	params      paramOption
}

func (b *baseInfo) Grouping() regalList {
	vHost := b.initialize()
	return b.calculate(vHost)
}

func (b *baseInfo) initialize() regalList {
	var l regalList

	for version, host := range b.versionHost {
		ipList := strings.Split(host, ",")
		l = append(l, []interface{}{version, ipList})
	}
	return l
}

func (b *baseInfo) calculate(vHost regalList) regalList {
	var baselist regalList
	for hostIndex, value := range vHost {
		hosts := value[1].([]string)[b.params.schedule:]
		baselist = append(baselist, []interface{}{value[0], [][]string{}})
		initHost := strings.Join(value[1].([]string)[:b.params.schedule], ", ")
		recursiveGrouping(hosts, baselist, initHost, b.params.combine, hostIndex)
	}
	return baselist
}

func Pop(items []string) (string, []string) {
	i := len(items) - 1
	popItem := items[i]
	items[i] = popItem
	items[len(items)-1] = ""
	items = items[:len(items)-1]
	return popItem, items
}

func recursiveGrouping(hosts []string, baselist regalList, init_host string, combine, hostindex int) int {
	var grouping func(int) int
	var init_n int
	var popitem string
	baselist[hostindex][1] = [][]string{{init_host}}

	grouping = func(init_n int) int {
		f_count := init_n + 1

		if len(hosts) == 0 {
			return 0
		}
		baselist[hostindex][1] = append(baselist[hostindex][1].([][]string), []string{})
		for i := 1; i <= combine; i++ {
			defer func() {
				if r := recover(); r != nil {
				}
			}()
			popitem, hosts = Pop(hosts)
			baselist[hostindex][1].([][]string)[f_count] = append(baselist[hostindex][1].([][]string)[f_count], popitem)
		}
		return grouping(f_count)
	}
	return grouping(init_n)

}
