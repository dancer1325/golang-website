# Go Tour

* Tour of Go
  * == Go programming language introduction
  * hosted [here](https://go.dev/tour/)

* Tour of Go's
  * app | this directory
  * [content](../_content/tour)

## Download/Install

* goal
  * install the tour -- from -- source

* allows
  * running offline the Go tour

* steps
  * [install Go](https://go.dev/doc/install)
  * `go install golang.org/x/website/tour@latest`
    * store a `tour` binary | your [GOPATH's `bin` directory](https://go.dev/cmd/go/#hdr-GOPATH_and_Modules)

## How to run locally?
* `go run .`
* | your browser,
  * opens http://localhost:3999

## Send Patches

* code changes -- via -- Gerrit 
* [how to contribute](https://go.dev/doc/contribute) 

## Report Issues

* if you want to report
  * [tour's code issues](https://github.com/golang/go/issues)
    * TODO:
  * [tour's content itself issues](https://github.com/golang/tour/issues)

## Deploying

* if a CL is reviewed & submitted -> the tour is AUTOMATICALLY deployed -- to -- App Engine
  * == -- as part of the -- main go.dev web site
* see [here](../cmd/golangorg/README.md#deploying-to-godev-and-golangorg)
