---
title: go fmt your code
date: 2013-01-23
by:
- Andrew Gerrand
tags:
- gofix
- gofmt
- technical
summary: How and why to format your Go code using gofmt.
---

## Introduction

[Gofmt](/cmd/gofmt/)

* := tool / 
  * AUTOMATICALLY formats Go source code
  * 's code is
    - easier to **write**
      - -> NEVER worry about MINOR formatting concerns 
    - easier to **read**
      - Reason: 🧠ALL code looks the same🧠
    - easier to **maintain**
      - Reason: 🧠diffs show ONLY the real changes🧠
    - **uncontroversial**
      - Reason: 🧠NEVER have a debate -- about -- spacing OR brace position🧠

## Format your code

* Go packages
  * 70%
    * are formatted -- according to -- gofmt's rules

* ways to format
  * `gofmt -w yourcode.go`
    * format your code -- via -- gofmt tool
  * `go fmt path/to/your/package`
    * format your code -- via -- [go fmt](/cmd/go/#hdr-Gofmt__reformat__package_sources) command

* plugins
  * [Vim plugin for Go](https://github.com/fatih/vim-go)
  * [go-mode.el](https://github.com/dominikh/go-mode.el)
    * == gofmt-before-save hook
  * [GoClipse](https://github.com/GoClipse/goclipse)
  * [GoSublime](https://github.com/DisposaBoy/GoSublime)
  * [misc/git/pre-commit script](https://github.com/golang/go/blob/release-branch.go1.1/misc/git/pre-commit)
  * [hgstyle plugin](https://bitbucket.org/fhs/hgstyle/overview)

## Mechanical source transformation

* machine-formatted code
  * 's mechanical transformation
    * | work with LARGE code bases, invaluable 
    * vs making wide-sweeping changes MANUALLY,
      * MORE comprehensive
      * less error-prune 
  * | working at scale,
    * NOT recommended to
      * make these kinds of changes MANUALLY

* `gofmt's -r`
  * the easiest way to -- MECHANICALLY manipulate -- Go code
  * == rewrite rule /
    ```
    pattern -> replacement
    # pattern          == Go expression
    # replacement      == Go expression
    ```
    * | `pattern`,
      * 1!-character lowercase identifiers == wildcards / match ARBITRARY sub-expressions
    * `replacement`
      * substitute the expressions

* [gofix](/cmd/fix/)
  * makes ARBITRARILY complex source transformations

* uses
  * | early days,
    * Reason: 🧠there were breaking changes | language & libraries🧠

* see [this article](/blog/introducing-gofix)
