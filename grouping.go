package regal

import (
	"fmt"
	"strings"
)

type baseInfo struct {
	versionHost map[string]string
	params      paramOption
}

func (b *baseInfo) Grouping() string {
	vHost := b.initialize()
	b.calculate(vHost)
	return ""
}

func (b *baseInfo) initialize() [][]interface{} {
	var l [][]interface{}

	for version, host := range b.versionHost {
		ipList := strings.Split(host, ",")
		l = append(l, []interface{}{version, ipList})

	}
	return l
}

func (b *baseInfo) calculate(vHost [][]interface{}) {
	var baselist [][]interface{}
	for hostIndex, value := range vHost {
		hosts := value[1].([]string)[b.params.schedule:]
		baselist = append(baselist, []interface{}{value[0], [][]string{}})
		initHost := strings.Join(value[1].([]string)[:b.params.schedule], ", ")
		fmt.Println(hosts, initHost, baselist)
		fmt.Println(recursiveGrouping(hosts, baselist, initHost, b.params.combine, hostIndex))
	}
	fmt.Println(baselist)
}

func pop(items []string) (string, []string) {
	length := len(items)
	popitem := items[length-1]
	items = items[:length-1]
	return popitem, items

}

func recursiveGrouping(hosts []string, baselist [][]interface{}, init_host string, combine, hostindex int) int {
	var grouping func(int) int
	var init_n int
	var popitem string
	baselist[0][1] = [][]string{{init_host}}

	grouping = func(init_n int) int {
		f_count := init_n + 1

		if len(hosts) == 0 {
			return 0
		}
		baselist[0][1] = append(baselist[0][1].([][]string), []string{})
		for i := 1; i <= combine; i++ {
			popitem, hosts = pop(hosts)
			baselist[0][1].([][]string)[f_count] = append(baselist[0][1].([][]string)[f_count], popitem)

		}
		return grouping(f_count)
	}
	return grouping(init_n)

}
