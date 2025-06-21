# Goal
* develop a Go package inside a module
* Check golang repo 'standardLibrary/cmd/go'

# Package
* == way to organize Go programs
* := collection of source files in the same directory which  -- Check in this directory how all contains `package main` --
  * 👁️ compiled together 👁️
  * functions & types & variables & constants in a source file 👁️— are visible to — rest of package’s files 👁️ -- TODO: Fix how --
* `import “ModulePathAndSubdirectoryWithinModule”` 
  * allows importing a package
* `go run main2.go`

# Module — “go.mod” —
* ≥ 1 module / repository
  * **Note:** 👁️Normally there is 1! module / repository located at root of the repository 👁️ -- Check in this project -- 
* := collection of related Go packages / released together
  * also contains package’s packages
* `module path`
  * `path`
    * 👁️ import path prefix 👁️ for all module’s packages -- Check 'https://pkg.go.dev/golang.org/x/website' for this project --
      * 👁️ @Standard Library ‘s packages do NOT contain the prefix 👁️
    * → indicates where to download it

# How to create a Go program?
* `mkdir HowToCreateGoProgram` 
* `go mod init example/user/hello` creates  the 'go.mod'
* Create a .go file
  * `package NameOfThePackage` — first statement —
    * `package main` — if you want executable commands —
* `go install NameOfTheModuleToCreate` — build & install the Go program — 
  * 👁️ == install a binary 👁️
  * if `GOBIN` is set → binary is installed under `GOBIN`
    * `go env` to print the Go environment variables
  * if `GOPATH` is set → binary is installed under `GOPATH`’s entry /bin
    * Problems
      * Problem1: Set -- GOPATH='/Users/dancer13/Library/Caches/go-build' -- but no bin found there
        * Attempt1: `go env -u GOPATH`
        * Attempt2: Open .bash_profile and switch to ''
        * Solution: TODO: 
  * else → binary is installed under '$HOME/go/bin'

## Import packages from your module
* Create the subdirectory '/morestrings' & add a file
* `cd morestrings` & `go build`
  * compiled package is saved in the local build cache
    * Problems: Where to find it?
      * Attempt1: `go env` & check GOCACHE & `cd $GOCACHE` 
      * Attempt2: `go env` & check GOCACHE & `cd $GOMODCACHE`
      * Solution: TODO:
* Package name can be used & install it -- `go install example/user/hello`
* `hello` runs the binary
  * Notes: Where is it placed ? -- TODO:

## Import packages from remote modules
* Add as common 'import' and by default it uses the 
* `go mod tidy` -- download & record & remove missing modules --

# Testing
* lightweight test framework
  * == `go test` + [testing](https://pkg.go.dev/testing@go1.22.1)
  * run each function
  * if the function calls a failure function → test is considered as failed
    * *Example of failure function:* `t.Error` OR `t.Fail`
* 👀how to create? 👀
  * file
    * name —  `*_test.go` —
    * with functions
      * name — `TestXXX` —
      * signature — `(t *testing.T)` —
* how to run?
  * `go test`
