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

English | [简体中文] | [Regal by Python]

## what's Regal-Go

For a simple example, let's say you need to do a staged rollout for a particular version  or several, which could be a bunch of server clusters, as shown in the following diagram:

<p align="center">
<img src="https://github.com/boylegu/regal-go/blob/main/image/fig01.png?raw=true">
</p>
</div>


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


