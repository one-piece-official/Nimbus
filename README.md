<p align="center">
  <img width="140"src="./logo.png">
</p>

<div align="center">
  <strong>
    收集一些常用的操作函数，辅助更快的完成开发工作，并减少重复代码。
  </strong>
</div>
<br />

<div align="center">
  <a href="https://app.dependabot.com/accounts/one-piece-official/repos/333034167">
    <img src="https://api.dependabot.com/badges/status?host=github&repo=one-piece-official/Nimbus&identifier=333034167" alt="Dependabot">
  </a>
  <img src="https://github.com/one-piece-official/Nimbus/workflows/Test/badge.svg" alt="Test">
  <img src="https://github.com/one-piece-official/Nimbus/workflows/Linter/badge.svg" alt="GolangCI Linter">
  <img src="https://github.com/one-piece-official/Nimbus/workflows/codecov/badge.svg" alt="codecov">
  <a href="https://pkg.go.dev/github.com/one-piece-official/Nimbus">
    <img src="https://img.shields.io/badge/godoc-ref-green.svg?style=flat" alt="GoDoc">
  </a>
  <a href="https://goreportcard.com/report/github.com/one-piece-official/Nimbus">
    <img src="https://goreportcard.com/badge/github.com/one-piece-official/Nimbus" alt="Go Report Card">
  </a>
  <a href="https://www.codefactor.io/repository/github/one-piece-official/nimbus"><img src="https://www.codefactor.io/repository/github/one-piece-official/nimbus/badge" alt="CodeFactor" /></a>
  <a href="https://github.com/one-piece-official/Nimbus/releases">
    <img src="https://img.shields.io/github/v/tag/one-piece-official/Nimbus.svg?label=release">
  </a>
</div>

## About Nimbus

参加魁地奇比赛，首先需要一把飞天扫帚，这也是巫师们最常使用的交通工具。基本分为横扫系列，彗星系列和光轮系列。罗恩曾经收到一支全新的横扫十一号做为礼物。马尔福家的扫帚是彗星一百二十号（Comet one twenty）。哈利·波特的飞天扫帚就是光轮2000（Nimbus2000）。

![footer](https://raw.githubusercontent.com/gobridge/about-us/master/gb_header.png)

## Overview

1. 以保证性能为前提，提供有效的方法；
2. 部分函数可以用来替换标准库函数，并与标准库函数保持一致的功能；
3. 部分函数仅会在一些特定的业务场景下有效。

### Standard

* 与标准库包名一致的使用 `ex` 前缀， 避免与标准库包冲突；
* 子包的 README.md 需要详述该包的用途和用例。

## Technologies

Todo.

## Getting Started

### Installation

```go
go get github.com/one-piece-official/Nimbus
```

### Features

Package `hash` functions.

| Function                                                     | Description  | #                                                            |
| ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ |
| [MD5](https://pkg.go.dev/github.com/one-piece-official/Nimbus/hash#MD5) | Fast hashing | [1138ac2](https://github.com/one-piece-official/Nimbus/commit/1138ac23a6e15cfd2ae58fabf20de573f49f6497) |
| [SHA1](https://pkg.go.dev/github.com/one-piece-official/Nimbus/hash#SHA1) | -            | -                                                            |


### Usage

Todo.

## Release History

* 0.2.1
  * ADD: 限量时间控制
* 0.2.0
  * ADD: ADD OS parser `setDefaultXYZ()`
* 0.1.1
  * ADD: Add new hash module

## Contributing to Nimbus
<!--- If your README is long or you have some specific process or steps you want contributors to follow, consider creating a separate CONTRIBUTING.md file--->
To contribute to Nimbus, follow these steps:

1. Fork this repository.
2. Clone it.
3. Push your job on your repository
4. When the job is done, submit a pull request, and that's it!

Alternatively see the GitHub documentation on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).

## Contributors ✨

Thanks to the following people who have contributed to this project:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/chenhang"><img src="https://avatars1.githubusercontent.com/u/3467833?v=4" width="80px;" alt="chenhang"/><br /><sub><b>Hang CHEN</b></sub></a></td>
    <td align="center"><a href="https://github.com/GiaoGiaoCat"><img src="https://avatars.githubusercontent.com/u/173622?v=4" width="80px;" alt="GiaoGiaoCat"/><br /><sub><b>GiaoGiaoCat</b></sub></a></td>
  </tr>
</table>

## Acknowledgements

* [go-extend](https://github.com/thinkeridea/go-extend)
* [go-funk](https://github.com/thoas/go-funk)
* [xstrings](https://github.com/huandu/xstrings)

## LICENSE

Personal uses are governed by the [MIT License](<https://github.com/one-piece-official/Nimbus/blob/main/LICENSE>). See LICENSE for more information.
