<div align="center">

<h1 style="border-bottom: none">
    <b style="color: #E940AF">Regal</b><b style="color: #66C9D6">-Go</b><br />
</h1>
<p>
用于"灰度发布"或 A/B Testing的智能分组引擎(Golang版)
</p>

[![go](https://img.shields.io/badge/Go-1.18+-66C9D6)]()
[![ver](https://img.shields.io/badge/version-1.0.0-66C9D6)]()
[![Maintainability](https://api.codeclimate.com/v1/badges/4c478e05a95251b6a818/maintainability)](https://codeclimate.com/github/boylegu/regal-go/maintainability)
[![go](https://img.shields.io/badge/license-MIT-E940AF)]()

<p align="center">
<img src="https://github.com/boylegu/regal-go/blob/main/image/regal-white.png?raw=true" width="330" height="300">
</p>
</div>

English | [简体中文] | [Python版本](https://github.com/boylegu/regal)

## Regal能做什么？
举个最简单的例子，比如需要针对一个版本进行灰度发布，而这一版本对应的可能是一大堆服务器集群， 如下图:

<p align="center">
<img src="https://github.com/boylegu/regal-go/blob/main/image/fig01_cn.png?raw=true">
</p>

就像图中描述的一样，无论你的服务器是多少，尤其很多中小型组织在进行灰度发布时，通常会面临分流策略在实际的技术或开发中如何去实现；

因此让Regal引擎直接介入，让它来根据你的策略进行动态地分组分流。 这里提供了两个参数：

- Combine

表示每组中的机器数量

- Schedule

作为A/B的第一组，默认为1。 可以通过使用'schedule'参数更改此行为。

## 功能

1. 提供分组策略，动态分流；
2. 支持多版本分组以及优先级可配置能力；

## 示例

详细示例可以查看[./example](./example).

### 示例1
```go
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
}
```
Output:

```shell
[root@gbe-pcx example]# go run main.go
[[app-test-version1.0 [[10.1.1.1] [10.1.1.2 10.1.1.3] [10.1.1.4 10.1.1.5]]]]
```

根据策略设置，会得到一个数据结构，这里可以观察一下：

<p align="center">
<img src="https://github.com/boylegu/regal-go/blob/main/image/fig02_cn.png?raw=true">
</p>

### Example-2

```go
	var example2 = [][]string{
		{"ver1", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5,10.1.1.6"},
		{"ver2", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
		{"ver3", "10.1.1.1,10.1.1.2,10.1.1.3,10.1.1.4,10.1.1.5"},
	}
	c2 := regal.RegalEngine(
		example2,
		regal.WithCombine(3),
		regal.WithSchedule(2),
		regal.WithPriorKey("ver2"), // Set priority
	)
	for _, v := range c2.Grouping() {
		fmt.Println(v)
	}

```

Output:

```shell
[root@gubaoer-pcx example]# go run main.go
[ver2 [[10.1.1.1, 10.1.1.2] [10.1.1.3 10.1.1.4 10.1.1.5]]]
[ver1 [[10.1.1.1, 10.1.1.2] [10.1.1.3 10.1.1.4 10.1.1.5] [10.1.1.6]]]
[ver3 [[10.1.1.1, 10.1.1.2] [10.1.1.3 10.1.1.4 10.1.1.5]]]
```

### Darwin's finches

<p align="let">
<img src="https://github.com/boylegu/regal-go/blob/main/image/b.jpg?raw=true" width="200" height="200">
</p>

人类的创造从来没有离开大自然带给我们的启发，而无论是灰度发布，还是A/B Testing，早在千年以前，大自然早有绝佳的解决方案。因此我以‘Darwin's finches’作为原型，向伟大的大自然和达尔文《物种起源》致敬。
> Author: Boyle Gu. Drawing with DeepAI in 2024.