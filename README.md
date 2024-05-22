<div align="center">

<h1 style="border-bottom: none">
    <b style="color: #E940AF">Regal</b><b style="color: #66C9D6">-Go</b><br />
</h1>
<p>
The smart grouping engine for A/B Testing or Gray release by Golang.
</p>

[![go](https://img.shields.io/badge/Go-1.18+-66C9D6)]()
[![ver](https://img.shields.io/badge/version-1.0.0-66C9D6)]()
[![Maintainability](https://api.codeclimate.com/v1/badges/4c478e05a95251b6a818/maintainability)](https://codeclimate.com/github/boylegu/regal-go/maintainability)
[![go](https://img.shields.io/badge/license-MIT-E940AF)]()

<p align="center">
<img src="https://github.com/boylegu/regal-go/blob/main/image/regal-white.png?raw=true" width="330" height="300">
</p>
</div>

English | [简体中文] | [Regal by Python](https://github.com/boylegu/regal)

## what's Regal-Go

For a simple example, let's say you need to do a staged rollout for a particular version  or several, which could be a bunch of server clusters, as shown in the following diagram:

<p align="center">
<img src="https://github.com/boylegu/regal-go/blob/main/image/fig01.png?raw=true">
</p>


The regal-go provides two policies:

- Combine

Number of machines in each group.

- Schedule

As the first group of A/B, the default is 1.
You can change this behavior by using the 'schedule' parameter.

>> See the example for more details.

## Feature

1. Provide A/B Test or Gray release policies and dynamic intelligent distribution;
2. Support multi-version grouping and priority;
3. Lightweight and scalable;

## Examples

See  [./example](./example) for example usage.

### Example-1
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
[root@gubaoer-pcx example]# go run main.go
[[app-test-version1.0 [[10.1.1.1] [10.1.1.2 10.1.1.3] [10.1.1.4 10.1.1.5]]]]
```
Based on policy, you will get a data structure. Let's take a look at it:

<p align="center">
<img src="https://github.com/boylegu/regal-go/blob/main/image/fig02.png?raw=true">
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

Human creation has never left the inspiration brought to us by nature, and whether it is gray release or A/B testing, nature had excellent solutions thousands of years ago.
Therefore, I use 'Darwin's finches' as the prototype to pay tribute to the great nature and Darwin's 《ORIGIN OF SPECIES》.

> Author: Boyle Gu. Drawing with DeepAI in 2024.